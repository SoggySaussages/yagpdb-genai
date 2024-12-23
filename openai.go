package genai

import (
	"context"
	"encoding/json"
	"html/template"
	"reflect"
	"slices"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/shared"
	"github.com/pkoukk/tiktoken-go"
	tiktoken_loader "github.com/pkoukk/tiktoken-go-loader"
)

type GenAIProviderOpenAI struct{}

func (p GenAIProviderOpenAI) ID() GenAIProviderID {
	return GenAIProviderOpenAIID
}

func (p GenAIProviderOpenAI) String() string {
	return "OpenAI"
}

func (p GenAIProviderOpenAI) DefaultModel() string {
	return openai.ChatModelGPT4oMini // cheapest model as of Dec 2024
}

var GenAIModelMapOpenAI = &GenAIProviderModelMap{
	"o1 Preview":    openai.ChatModelO1Preview,
	"o1 Mini":       openai.ChatModelO1Mini,
	"GPT 4o":        openai.ChatModelGPT4o,
	"ChatGPT 4o":    openai.ChatModelChatgpt4oLatest,
	"GPT 4o Mini":   openai.ChatModelGPT4oMini,
	"GPT 4 Turbo":   openai.ChatModelGPT4Turbo,
	"GPT 4":         openai.ChatModelGPT4,
	"GPT 3.5 Turbo": openai.ChatModelGPT3_5Turbo,
}

func (p GenAIProviderOpenAI) ModelMap() *GenAIProviderModelMap {
	return GenAIModelMapOpenAI
}

func (p GenAIProviderOpenAI) KeyRequired() bool {
	return true
}

var globalConfigGenAI = &GenAIProviderGlobalConfig{}

func (p GenAIProviderOpenAI) GlobalConfig() *GenAIProviderGlobalConfig {
	return globalConfigGenAI
}

// ~ accurate for English text as of Dec 2024
const CharacterCountToTokenRatioOpenAI = 4 / 1

func (p GenAIProviderOpenAI) CharacterTokenRatio() int {
	return CharacterCountToTokenRatioOpenAI
}

var tiktokenBPELoaderSet bool

func (p GenAIProviderOpenAI) EstimateTokens(model, combinedInput string, maxTokens int64) (inputEstimatedTokens, outputMaxTokens int64) {
	if !tiktokenBPELoaderSet {
		tiktoken.SetBpeLoader(tiktoken_loader.NewOfflineLoader())
		tiktokenBPELoaderSet = true
	}

	inputEstimatedTokens = int64(len(combinedInput) / CharacterCountToTokenRatioOpenAI)
	outputMaxTokens = maxTokens - (inputEstimatedTokens / 4)
	return
}

func (p GenAIProviderOpenAI) ValidateAPIToken(key string) error {
	// make a really cheap (%0.02 of a cent) call to test the key
	client := openai.NewClient(option.WithAPIKey(key))
	_, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Messages:            openai.F([]openai.ChatCompletionMessageParamUnion{openai.UserMessage("1")}),
		Model:               openai.F(p.DefaultModel()),
		MaxCompletionTokens: openai.Int(1),
	})
	return err
}

var ModelsNotSupportingSystemRoleMessages = []string{openai.ChatModelO1Mini, openai.ChatModelO1Preview}

func (p GenAIProviderOpenAI) ComplexCompletion(model, key string, input *GenAIInput) (*GenAIResponse, *GenAIResponseUsage, error) {
	messages := []openai.ChatCompletionMessageParamUnion{}

	if slices.Contains(ModelsNotSupportingSystemRoleMessages, model) {
		messages = append(messages, openai.UserMessage(input.BotSystemMessage))
	} else {
		messages = append(messages, openai.SystemMessage(input.BotSystemMessage))
	}

	if input.SystemMessage != "" {
		if slices.Contains(ModelsNotSupportingSystemRoleMessages, model) {
			messages = append(messages, openai.UserMessage(input.SystemMessage))
		} else {
			messages = append(messages, openai.SystemMessage(input.SystemMessage))
		}
	}

	if input.UserMessage != "" {
		messages = append(messages, openai.UserMessage(input.UserMessage))
	}

	var tools []openai.ChatCompletionToolParam

	if input.Functions != nil {
		for _, fn := range *input.Functions {
			properties := make(map[string]interface{}, 0)
			for argName, argType := range fn.Arguments {
				properties[argName] = map[string]string{
					"type": argType,
				}
			}

			tools = append(tools, openai.ChatCompletionToolParam{
				Type: openai.F(openai.ChatCompletionToolTypeFunction),
				Function: openai.F(openai.FunctionDefinitionParam{
					Name:        openai.String(fn.Name),
					Description: openai.String(fn.Description),
					Parameters: openai.F(openai.FunctionParameters{
						"type":       "object",
						"properties": properties,
					}),
				}),
			})
		}
	}

	requestParams := openai.ChatCompletionNewParams{Model: openai.F(model), Messages: openai.F([]openai.ChatCompletionMessageParamUnion{openai.UserMessage("Please begin.")}), MaxCompletionTokens: openai.Int(input.MaxTokens), Temperature: openai.Float(1.1), PresencePenalty: openai.Float(0.1)}

	if len(messages) > 0 {
		requestParams.Messages = openai.F(messages)
	}

	if len(tools) > 0 {
		requestParams.Tools = openai.F(tools)
	}

	client := openai.NewClient(option.WithAPIKey(key))

	usage := &GenAIResponseUsage{}

	chatCompletion, err := client.Chat.Completions.New(context.Background(), requestParams)
	if chatCompletion != nil && chatCompletion.Usage.PromptTokens != 0 || chatCompletion.Usage.CompletionTokens != 0 {
		usage.InputTokens = chatCompletion.Usage.PromptTokens
		usage.OutputTokens = chatCompletion.Usage.CompletionTokens
	}
	if err != nil {
		return nil, usage, err
	}

	choice := chatCompletion.Choices[0]
	content := choice.Message.Content
	if choice.Message.Refusal != "" {
		content = choice.Message.Refusal
	}

	var functionResponse []GenAIFunctionResponse
	if len(choice.Message.ToolCalls) > 0 {
		currentFunc := GenAIFunctionResponse{}
		functionCall := choice.Message.ToolCalls[0].Function
		currentFunc.Name = functionCall.Name
		json.Unmarshal([]byte(functionCall.Arguments), &currentFunc.Arguments)
		functionResponse = append(functionResponse, currentFunc)
	}

	return &GenAIResponse{
		Content:   content,
		Functions: &functionResponse,
	}, usage, nil
}

func (p GenAIProviderOpenAI) ModerateMessage(model, key string, message string) (*GenAIModerationCategoryProbability, *GenAIResponseUsage, error) {
	moderationParams := openai.ModerationNewParams{
		Input: openai.F[openai.ModerationNewParamsInputUnion](shared.UnionString(message)),
		Model: openai.F(openai.ModerationModelOmniModerationLatest),
	}

	client := openai.NewClient(option.WithAPIKey(key))

	inputUse, _ := p.EstimateTokens(openai.ModerationModelOmniModerationLatest, message, 0)
	usage := &GenAIResponseUsage{InputTokens: inputUse}

	moderation, err := client.Moderations.New(context.Background(), moderationParams)
	if err != nil {
		return nil, usage, err
	}

	response := GenAIModerationCategoryProbability{}

	catScores := reflect.ValueOf(moderation.Results[0].CategoryScores)
	typed := catScores.Type()
	availableCategories := []string{}
	for _, c := range GenAIModerationCategories {
		availableCategories = append(availableCategories, strings.ReplaceAll(c, " ", ""))
	}

	for i := 0; i < catScores.NumField(); i++ {
		category := typed.Field(i).Name
		if !slices.Contains(availableCategories, strings.ToLower(category)) {
			continue
		}

		score := catScores.Field(i).Float()
		response[category] = score
	}

	return &response, usage, nil
}

var GenAIProviderOpenAIWebData = &GenAIProviderWebDescriptions{
	ObtainingAPIKeyInstructions: template.HTML(`Step one: Create an account.
	<br>
	Visit <a href="https://platform.openai.com/docs/guides/production-best-practices/api-keys#setting-up-your-organization">OpenAI's website</a> to do this.
	<br>
	<br>
	Step two: Set up payment method.
	<br>
	You must set up a payment method in order to make requests to OpenAI. Do so on <a href="https://platform.openai.com/settings/organization/billing/overview">OpenAI's API dashboard</a>.
	<br>
	<br>
	Step three: Set a Budget Limit.
	<br>
	You must set a monthly budget limit within reason to prevent yourself from going into credit debt with OpenAI. Do so on <a href="https://platform.openai.com/settings/organization/limits">OpenAI's API dashboard</a>.
	<br>
	<br>
	Step four: Create an API key.
	<br>
	Create an API key on <a href="https://platform.openai.com/api-keys">OpenAI's Dashboard</a>. Set the mode to <strong>restricted</strong>, set every permission to <strong>None</strong>, and then set the "Model capabilities" permission to <strong>Write</strong>.
	<br>
	<br>
	Step five: Copy the API key to YAGPDB.
	<br>
	Click copy, then paste the new API key into the "API Key" field on this page.`),
	ModelDescriptionsURL: "https://platform.openai.com/docs/models",
	ModelForModeration:   "omni-moderation-latest",
	PlaygroundURL:        "https://platform.openai.com/playground/chat?models=gpt-4o-mini",
}

func (p GenAIProviderOpenAI) WebData() *GenAIProviderWebDescriptions {
	return GenAIProviderOpenAIWebData
}
