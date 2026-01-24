package handlers

import (
	"net/http"

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
