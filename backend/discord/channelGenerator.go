package discord

import (
	"backend/pocketbase"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) handleUserJoinedChannelGenerator(discord *discordgo.Session, update *discordgo.VoiceStateUpdate) {
	joinedChannel := update.ChannelID
	if joinedChannel == "" {
		// User left channel; No actions needed
		return
	}
	config, err := b.PocketBaseClient.ChannelGeneratorConfigForChannelID(joinedChannel)
	if err != nil {
		b.Logger.Error("Error retrieving channel config", err)
		return
	}
	if config == nil {
		// Channel has no config; No actions needed
		return
	}

	b.newVoiceChannelGeneratorChannelFromConfig(discord, config, update.UserID)
}

func (b *Bot) handleUserLeftChannelGeneratorChannel(discord *discordgo.Session, update *discordgo.VoiceStateUpdate) {
	if update.BeforeUpdate == nil {
		b.Logger.Info("No info about previous channel membership")
		return
	}
	leftChannel := update.BeforeUpdate.ChannelID
	channel, err := b.PocketBaseClient.OwnedChannelGeneratorChannel(leftChannel, update.UserID)
	if err != nil {
		b.Logger.Error("Failed to get owned channel", err)
	}

	if channel == nil {
		return
	}

	users, err := b.UsersInVoiceChannel(discord, update.GuildID, leftChannel)
	if err != nil {
		b.Logger.Error("Failed to get users in voice channel")
		return
	}
	if len(users) == 0 {
		deletedChannel, err := discord.ChannelDelete(leftChannel)
		if err != nil {
			b.Logger.Error("Failed to delete channel: %q", deletedChannel)
			return
		}
		err = b.PocketBaseClient.RemoveChannelGeneratorChannel(leftChannel)
		if err != nil {
			b.Logger.Error("Failed to remove record for channel: %q", err)
		}
		return
	}

	config := channel.ExpandedOne("parentConfig")
	channelName := channel.GetName()

	nextOwner := users[0]
	switch config.GetString("nameFormat") {
	case "index":
		break
	case "owner":
		newOwnerUser, err := discord.GuildMember(update.GuildID, nextOwner)
		if err != nil {
			b.Logger.Error("Failed to get new owner info")
			return
		}
		channelName = fmt.Sprintf("%s's Channel", newOwnerUser.DisplayName())
	}

	if channelName != channel.GetName() {
		_, err := discord.ChannelEdit(leftChannel, &discordgo.ChannelEdit{Name: channelName})
		if err != nil {
			b.Logger.Error("Failed to rename channel: %q", err)
		}
	}

	err = b.PocketBaseClient.TransferChannelGeneratorChannelOwner(channel, nextOwner, channelName)
	if err != nil {
		b.Logger.Error("Failed to transfer channel record to new owner: %q", err)
	}
}

func (b *Bot) newVoiceChannelGeneratorChannelFromConfig(discord *discordgo.Session, config *pocketbase.VoiceChannelGeneratorConfig, ownerID string) {
	channelIndex, err := b.PocketBaseClient.NextFreeChannelGeneratorIndex(config.GetChannelID())
	if err != nil {
		b.Logger.Error("Error getting next free channel index")
		return
	}
	b.Logger.Debug(fmt.Sprintf("vals %+v || %d", config, channelIndex))
	channelName := ""
	switch config.GetNameFormat() {
	case "index":
		channelName = fmt.Sprintf("Channel %d", channelIndex)
	case "owner":
		newOwnerUser, err := discord.GuildMember(config.GetServerID(), ownerID)
		if err != nil {
			b.Logger.Error("Failed to get new owner info")
			return
		}
		channelName = fmt.Sprintf("%s's Channel", newOwnerUser.DisplayName())
	}

	parentChannel, err := discord.Channel(config.GetChannelID())
	if err != nil {
		b.Logger.Error(fmt.Sprintf("Error getting parent channel structure: %q", err))
		return
	}

	newChannelData := discordgo.GuildChannelCreateData{
		Name:                 channelName,
		Type:                 discordgo.ChannelTypeGuildVoice,
		Topic:                "",
		Bitrate:              parentChannel.Bitrate,
		UserLimit:            parentChannel.UserLimit,
		RateLimitPerUser:     parentChannel.RateLimitPerUser,
		Position:             parentChannel.Position,
		PermissionOverwrites: parentChannel.PermissionOverwrites,
		ParentID:             parentChannel.ParentID,
		NSFW:                 parentChannel.NSFW,
	}

	channel, err := discord.GuildChannelCreateComplex(config.GetServerID(), newChannelData)
	if err != nil {
		b.Logger.Error("Failed to create new channel:", err)
		return
	}

	err = b.PocketBaseClient.SaveChannelGeneratorChannel(channel.ID, ownerID, channelIndex, config.Id, channelName)
	if err != nil {
		b.Logger.Error("Failed to persist channel:", err)
	}

	err = discord.GuildMemberMove(config.GetServerID(), ownerID, &channel.ID)
	if err != nil {
		b.Logger.Error("Failed to move user to new channel", err)
	}

	return
}
