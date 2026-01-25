package handlers

import (
	"net/http"

	"github.com/chrishultin/SpediBot/backend/discord"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

// get configs for a specific server
type GetChannelGeneratorConfigsResponse struct {
	Configs []ChannelGeneratorConfig `json:"configs"`
}

type ChannelGeneratorConfig struct {
	ID           string `json:"id"`
	ServerID     string `json:"serverId"`
	ChannelID    string `json:"channelId"`
	NameFormat   string `json:"nameFormat"`
	EnableRename bool   `json:"enableRename"`
}

func HandleGetChannelGeneratorConfigs(e *core.RequestEvent) error {
	configs, err := e.App.FindRecordsByFilter("channel_generator_configs",
		"serverID = {:serverID}",
		"", 0, 0, dbx.Params{"serverID": e.Request.PathValue("serverID")})
	if err != nil {
		return err
	}

	output := make([]ChannelGeneratorConfig, len(configs))
	for i, config := range configs {
		output[i] = ChannelGeneratorConfig{
			ID:           config.Id,
			ServerID:     config.GetString("serverID"),
			ChannelID:    config.GetString("channelID"),
			NameFormat:   config.GetString("nameFormat"),
			EnableRename: config.GetBool("enableRename"),
		}
	}
	return e.JSON(http.StatusOK, GetChannelGeneratorConfigsResponse{Configs: output})
}

type UpdateChannelGeneratorConfigRequest struct {
	ID           string `json:"id"`
	ChannelID    string `json:"channelId"`
	NameFormat   string `json:"nameFormat"`
	EnableRename bool   `json:"enableRename"`
}

type UpdateChannelGeneratorConfigResponse struct {
	Config ChannelGeneratorConfig `json:"config"`
}

func HandleUpdateChannelGeneratorConfig(discordBot *discord.Bot) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		serverID := e.Request.PathValue("serverID")

		request := UpdateChannelGeneratorConfigRequest{}

		if err := e.BindBody(&request); err != nil {
			return err
		}

		if !discordBot.UserIsAdmin(e.Auth.GetString("discordID"), serverID) {
			return e.UnauthorizedError("User cannot configure server %q", serverID)
		}

		configurationsCollection, err := e.App.FindCollectionByNameOrId("channel_generator_configs")
		if err != nil {
			return err
		}

		var config *core.Record

		if request.ID != "" {
			config, err = e.App.FindRecordById("channel_generator_configs", request.ID)
			if err != nil {
				return err
			}
		} else {
			config = core.NewRecord(configurationsCollection)
			config.Set("serverID", serverID)
		}

		config.Set("nameFormat", request.NameFormat)
		config.Set("enableRename", request.EnableRename)
		config.Set("channelID", request.ChannelID)

		err = e.App.Save(config)
		if err != nil {
			return err
		}
		return e.JSON(http.StatusOK, UpdateChannelGeneratorConfigResponse{Config: ChannelGeneratorConfig{
			ID:           config.Id,
			ServerID:     config.GetString("serverID"),
			ChannelID:    config.GetString("channelID"),
			NameFormat:   config.GetString("nameFormat"),
			EnableRename: config.GetBool("enableRename"),
		}})
	}
}
