package genai

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/SoggySaussages/yagpdb-genai/models"
	"github.com/botlabs-gg/yagpdb/v2/lib/dstate"
	"github.com/botlabs-gg/yagpdb/v2/premium"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type GenAIProviderGeneric struct {
	provider GenAIProvider
}

func (p GenAIProviderGeneric) ID() GenAIProviderID {
	return p.provider.ID()
}

func (p GenAIProviderGeneric) String() string {
	return p.provider.String()
}

func (p GenAIProviderGeneric) DefaultModel() string {
	if p.GlobalConfig().Model != "" {
		return p.GlobalConfig().Model
	}
	var backupModel string
	for _, v := range *p.ModelMap() {
		if backupModel != "" {
			backupModel = v
		}
		if v == p.provider.DefaultModel() {
			return v
		}
	}

	// backup model is the first model allowed by config or an empty string if
	// no models are allowed by config. we only reach this if the hard coded
	// default model is not permitted by config AND there is no global override
	// for the model for this provider
	return backupModel
}

func (p GenAIProviderGeneric) ModelMap() *GenAIProviderModelMap {
	return p.provider.ModelMap()
}

func (p GenAIProviderGeneric) KeyRequired() bool {
	return p.provider.KeyRequired()
}

func (p GenAIProviderGeneric) GlobalConfig() *GenAIProviderGlobalConfig {
	return p.provider.GlobalConfig()
}

func (p GenAIProviderGeneric) CharacterTokenRatio() int {
	return p.provider.CharacterTokenRatio()
}

func (p GenAIProviderGeneric) EstimateTokens(model, combinedInput string, maxTokens int64) (inputEstimatedTokens, outputMaxTokens int64) {
	return p.provider.EstimateTokens(model, combinedInput, maxTokens)
}

func (p GenAIProviderGeneric) ValidateAPIToken(gs *dstate.GuildState, token string) error {
	_, key, err := getAPIToken(gs)
	if err != nil {
		return err
	}
	return p.provider.ValidateAPIToken(key)
}

func (p GenAIProviderGeneric) BasicCompletion(gs *dstate.GuildState, systemMsg, userMsg string, maxTokens int64, nsfw bool) (*GenAIResponse, *GenAIResponseUsage, error) {
	input := &GenAIInput{BotSystemMessage: BotSystemMessagePromptGeneric + BotSystemMessagePromptAppendSingleResponseContext, SystemMessage: systemMsg, UserMessage: userMsg, MaxTokens: maxTokens}
	if nsfw {
		input.BotSystemMessage = fmt.Sprintf("%s\n%s", input.BotSystemMessage, BotSystemMessagePromptAppendNSFW)
	} else {
		input.BotSystemMessage = fmt.Sprintf("%s\n%s", input.BotSystemMessage, BotSystemMessagePromptAppendNonNSFW)
	}
	return p.ComplexCompletion(gs, input)
}

func (p GenAIProviderGeneric) determineModel(conf *models.GenaiConfig) (model string) {
	model = conf.Model
	if p.GlobalConfig().Model != "" {
		model = p.GlobalConfig().Model
	}
	for _, v := range *p.ModelMap() {
		if v == model {
			return v
		}
	}
	return p.DefaultModel()
}

func (p GenAIProviderGeneric) calcInputMax(conf *models.GenaiConfig, input *GenAIInput) bool {
	combinedInput := input.BotSystemMessage + input.SystemMessage + input.UserMessage
	if input.Functions != nil && len(*input.Functions) > 0 {
		b, _ := json.Marshal(*input.Functions)
		combinedInput += string(b)
	}
	maxMonth := conf.MaxTokens
	if len(p.GlobalConfig().Key) > 0 {
		maxMonth = int64(confMaxMonthlyTokens.GetInt())
		if prem, _ := premium.IsGuildPremium(conf.GuildID); prem {
			maxMonth = int64(confMaxMonthlyTokensPremium.GetInt())
		}
	}
	calculatedMax := input.MaxTokens
	if maxMonth <= 0 {
		availableTokens := maxMonth - conf.MonthTokenUsageToDate
		calculatedMax = int64(math.Min(float64(input.MaxTokens), float64(availableTokens)))
	}

	_, maxOut := p.EstimateTokens(p.determineModel(conf), combinedInput, calculatedMax)
	input.MaxTokens = maxOut
	return maxOut > 0
}

var ErrTokenLimitReached = errors.New("not enough tokens in quota")

func (p GenAIProviderGeneric) updateUsage(conf *models.GenaiConfig, usage *GenAIResponseUsage) {
	whitelist := []string{"month_token_usage_to_date"}
	oneMonthElapsed := conf.TokenUsageLastReset.UTC().Add(time.Hour * 24 * 30).Before(time.Now().UTC())
	if conf.TokenUsageLastReset.UTC().Month() != time.Now().UTC().Month() || oneMonthElapsed {
		conf.MonthTokenUsageToDate = 0
		conf.TokenUsageLastReset = time.Now()
		whitelist = append(whitelist, "token_usage_last_reset")
	}
	conf.MonthTokenUsageToDate += usage.InputTokens/4 + usage.OutputTokens
	_, err := conf.UpdateG(context.Background(), boil.Whitelist(whitelist...))
	if err != nil {
		logger.WithError(err).WithField("guild", conf.GuildID).Error("URGENT: failed updating token counter, overages may occur")
	}
}

func (p GenAIProviderGeneric) ComplexCompletion(gs *dstate.GuildState, input *GenAIInput) (*GenAIResponse, *GenAIResponseUsage, error) {
	conf, key, err := getAPIToken(gs)
	if err != nil {
		if err == ErrorNoAPIKey {
			return &GenAIResponse{Content: "Please set your API key on the dashboard to use Generative AI."}, &GenAIResponseUsage{}, nil
		}
		if err == ErrorAPIKeyInvalid {
			return &GenAIResponse{Content: err.Error()}, &GenAIResponseUsage{}, nil
		}
		return nil, nil, err
	}

	if !p.calcInputMax(conf, input) {
		return &GenAIResponse{}, &GenAIResponseUsage{}, ErrTokenLimitReached
	}

	r, usage, err := p.provider.ComplexCompletion(p.determineModel(conf), key, input)
	if usage == nil || usage.InputTokens == 0 && usage.OutputTokens == 0 {
		return r, usage, err
	}

	p.updateUsage(conf, usage)
	return r, usage, err
}

func (p GenAIProviderGeneric) ModerateMessage(gs *dstate.GuildState, message string) (*GenAIModerationCategoryProbability, *GenAIResponseUsage, error) {
	conf, key, err := getAPIToken(gs)
	if err != nil {
		return &GenAIModerationCategoryProbability{}, nil, nil
	}

	fakeInput := CustomModerateFunction
	fakeInput.UserMessage = message
	if !p.calcInputMax(conf, &fakeInput) {
		return &GenAIModerationCategoryProbability{}, &GenAIResponseUsage{}, nil
	}

	r, usage, err := p.provider.ModerateMessage(p.determineModel(conf), key, message)
	if usage == nil || usage.InputTokens == 0 && usage.OutputTokens == 0 {
		return r, usage, nil
	}

	p.updateUsage(conf, usage)
	return r, usage, nil
}

func (p GenAIProviderGeneric) WebData() *GenAIProviderWebDescriptions {
	return p.provider.WebData()
}
