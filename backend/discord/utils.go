package discord

import (
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) UsersInVoiceChannel(discord *discordgo.Session, serverID, channelID string) ([]string, error) {
	var members []string
	guild, err := discord.State.Guild(serverID)
	if err != nil {
		return members, err
	}
	for _, guildMemberState := range guild.VoiceStates {
		if guildMemberState.ChannelID == channelID {
			members = append(members, guildMemberState.UserID)
		}
	}
	return members, nil
}
