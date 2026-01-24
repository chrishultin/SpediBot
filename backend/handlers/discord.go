package handlers

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/chrishultin/SpediBot/backend/discord"
	"github.com/pocketbase/pocketbase/core"
)

type GetAdminServersResponse struct {
	Servers []GetAdminServersServer `json:"servers"`
}

type GetAdminServersServer struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

func GetServers(discordBot *discord.Bot) func(e *core.RequestEvent) error {
	return func(event *core.RequestEvent) error {
		userID := event.Auth.GetString("discordID")
		servers := discordBot.UserIsAdminForServers(userID)

		resp := GetAdminServersResponse{
			Servers: make([]GetAdminServersServer, len(servers)),
		}
		for i, server := range servers {
			resp.Servers[i] = GetAdminServersServer{
				Name:        server.Name,
				ID:          server.ID,
				Description: server.Description,
				Icon:        server.Icon,
			}
		}

		return event.JSON(http.StatusOK, resp)
	}
}

type GetChannelsForServerResponse struct {
	Channels []GetChannelsForServerChannel `json:"channels"`
}

type GetChannelsForServerChannel struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	ServerID string `json:"server_id"`
}

func GetChannelsForServer(discordBot *discord.Bot) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		serverID := e.Request.PathValue("serverID")
		if serverID == "" {
			return e.Error(http.StatusBadRequest, "serverID is required", nil)
		}
		channels, err := discordBot.ChannelsForServer(serverID)
		if err != nil {
			return e.Error(http.StatusBadRequest, "could not get channels for server: "+serverID, err)
		}

		output := make([]GetChannelsForServerChannel, 0)
		for _, channel := range channels {
			if channel.Type != discordgo.ChannelTypeGuildVoice {
				continue
			}
			output = append(output, GetChannelsForServerChannel{
				Name:     channel.Name,
				ID:       channel.ID,
				ServerID: channel.GuildID,
			})
		}
		return e.JSON(http.StatusOK, GetChannelsForServerResponse{Channels: output})
	}
}
