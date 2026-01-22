package handlers

import (
	"net/http"

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
	return nil
}
