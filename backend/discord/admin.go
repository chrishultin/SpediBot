package discord

import "github.com/bwmarrin/discordgo"

type server struct {
	Name        string
	ID          string
	Description string
	Icon        string
}

func (b *Bot) UserIsAdminForServers(userID string) []server {
	var output []server

	for _, guild := range b.Session.State.Guilds {
		if b.UserIsAdmin(userID, guild.ID) {
			output = append(output, server{Name: guild.Name, ID: guild.ID, Description: guild.Description, Icon: guild.IconURL("32")})
		}
	}

	return output
}

func (b *Bot) UserIsAdmin(userID string, serverID string) bool {
	guild, err := b.Session.Guild(serverID)
	if err != nil {
		return false
	}
	if guild.OwnerID == userID {
		return true
	}

	guildMember, err := b.Session.GuildMember(guild.ID, userID)
	if err != nil {
		return false
	}

	for _, roleID := range guildMember.Roles {
		role, err := b.Session.State.Role(guild.ID, roleID)
		if err != nil {
			continue
		}
		if role.Permissions&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator {
			return true
		}
	}

	return false
}
