package discord

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	pocketbaseint "github.com/chrishultin/SpediBot/backend/pocketbase"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	AppID string
	Token string

	PocketBaseClient *pocketbaseint.Client
	Logger           *slog.Logger
	Session          *discordgo.Session
}

func (b *Bot) Serve() error {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", b.Token))
	if err != nil {
		return err
	}

	session.Identify.Intents = discordgo.IntentsAll
	session.StateEnabled = true

	session.AddHandler(b.handleUserJoinedChannelGenerator)
	session.AddHandler(b.handleUserLeftChannelGeneratorChannel)

	err = session.Open()
	if err != nil {
		panic(err)
	}
	b.Session = session

	defer func(s *discordgo.Session) {
		err := s.Close()
		if err != nil {
			b.Logger.Error("Discord bot crashed with error", err)
			panic(err)
		}
	}(session)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	<-stop
	fmt.Println("Bot has terminated")
	return nil
}
