package main

import (
	"backend/discord"
	"log"
	"os"

	pocketbaseint "backend/pocketbase"

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

	bot := discord.Bot{
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

		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
