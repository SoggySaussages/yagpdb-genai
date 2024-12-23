package genai

import (
	"context"
	"strconv"
	"strings"

	"github.com/botlabs-gg/yagpdb/v2/automod"
	"github.com/botlabs-gg/yagpdb/v2/bot"
	"github.com/botlabs-gg/yagpdb/v2/lib/discordgo"
	"github.com/botlabs-gg/yagpdb/v2/lib/dstate"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type GenAIAutomodTriggerData struct {
	Threshold  int
	MaxTokens  int
	Categories []string
}

var _ automod.MessageTrigger = (*GenAIAutomodTrigger)(nil)

type GenAIAutomodTrigger struct{}

func (mc *GenAIAutomodTrigger) Kind() automod.RulePartType {
	return automod.RulePartTrigger
}

func (mc *GenAIAutomodTrigger) DataType() interface{} {
	return &GenAIAutomodTriggerData{}
}

func (mc *GenAIAutomodTrigger) Name() string {
	return "Generative AI trigger"
}

func (mc *GenAIAutomodTrigger) Description() string {
	return "Triggers on messages AI marks abusive. Requires GenAI to be enabled."
}

func (mc *GenAIAutomodTrigger) UserSettings() []*automod.SettingDef {
	settings := []*automod.SettingDef{
		{
			Name:    "Certainty Threshold (%)",
			Key:     "Threshold",
			Kind:    automod.SettingTypeInt,
			Default: 60,
			Min:     1,
			Max:     99,
		},
		{
			Name:    "Max Tokens per Message",
			Key:     "MaxTokens",
			Kind:    automod.SettingTypeInt,
			Default: 512,
			Min:     1,
			Max:     512,
		},
	}
	catSettings := &automod.SettingDef{
		Name: "Categories to Trigger",
		Key:  "Categories",
		Kind: automod.SettingTypeMultiOptionsCustom,
	}
	for i, s := range GenAIModerationCategories {
		catSettings.Options = append(catSettings.Options, automod.SettingTypeOptionsCustomOption{
			Name:  GenAIModerationCategoriesFormatted[i],
			Value: strings.ReplaceAll(s, " ", "-"),
		})
	}
	return append(settings, catSettings)
}

func (mc *GenAIAutomodTrigger) CheckMessage(triggerCtx *automod.TriggerContext, cs *dstate.ChannelState, m *discordgo.Message) (bool, error) {
	dataCast := triggerCtx.Data.(*GenAIAutomodTriggerData)
	config, err := GetConfig(cs.GuildID)
	if err != nil {
		return false, err
	}

	if !config.Enabled || len(config.Key) == 0 {
		return false, nil
	}

	provider := GenAIProviderFromID(config.Provider)

	if config.OwnerIDAutomodNotified != triggerCtx.GS.OwnerID {
		sent := true
		alertMsg := "Heads up: YAGPDB just used Generative AI to analyze its first message on your server for abuse. It's important to notify all of your users (preferably in a join DM as well as a prominent channel message in your server) that messages they send may be forwarded to " + provider.String() + " for AI analysis."
		err := bot.SendDM(triggerCtx.GS.OwnerID, alertMsg)
		if err != nil {
			permsOk, _, err := bot.SendMessage(cs.GuildID, cs.ID, "<@"+strconv.FormatInt(triggerCtx.GS.OwnerID, 10)+"> "+alertMsg)
			if !permsOk {
				permsOk, _, err = bot.SendMessage(cs.GuildID, cs.GuildID, "<@"+strconv.FormatInt(triggerCtx.GS.OwnerID, 10)+"> "+alertMsg)
			}
			if !permsOk || err != nil {
				sent = false
			}
		}
		if sent {
			config.OwnerIDAutomodNotified = triggerCtx.GS.OwnerID
			config.UpdateG(context.Background(), boil.Whitelist("owner_id_automod_notified"))
		}
	}

	g := bot.State.GetGuild(cs.GuildID)
	if g == nil {
		return false, err
	}

	content := m.Content
	maxContentLength := dataCast.MaxTokens * provider.CharacterTokenRatio()
	if len(content) > maxContentLength {
		content = content[:maxContentLength]
	}

	categories, _, err := provider.ModerateMessage(&dstate.GuildState{ID: g.ID, OwnerID: g.OwnerID}, content)
	if err != nil {
		logger.WithError(err).Error("GenAI Automod trigger API error")
		return false, nil
	}
	for _, c := range dataCast.Categories {
		pascalCaseCatName := strings.ReplaceAll(strings.Title(strings.ReplaceAll(c, "-", " ")), " ", "")
		if (*categories)[pascalCaseCatName]*100 > float64(dataCast.Threshold) {
			return true, nil
		}
	}
	return false, nil
}
