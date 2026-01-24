package main

import (
	"log"
	"os"

	"github.com/chrishultin/SpediBot/backend/discord"
	"github.com/chrishultin/SpediBot/backend/handlers"
	"github.com/pocketbase/pocketbase/apis"

	pocketbaseint "github.com/chrishultin/SpediBot/backend/pocketbase"
	frontend "github.com/chrishultin/SpediBot/ui-embed"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

var DISCORD_TOKEN = os.Getenv("DISCORD_BOT_TOKEN")
var APP_ID = os.Getenv("DISCORD_APP_ID")
var DEV_MODE = os.Getenv("DEV_MODE") != ""

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: DEV_MODE,
	})

	bot := &discord.Bot{
		AppID:            APP_ID,
		Token:            DISCORD_TOKEN,
		PocketBaseClient: &pocketbaseint.Client{PocketBase: app},
		Logger:           app.Logger(),
	}

	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		go func() {
			err := bot.Serve()
			if err != nil {
				panic(err)
			}
		}()

		e.Router.GET("/{path...}", apis.Static(frontend.EmbeddedUI, true))
		e.Router.GET("/api/custom/servers", handlers.GetServers(bot))
		e.Router.GET("/api/custom/servers/{serverID}/channels", handlers.GetChannelsForServer(bot))

		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
