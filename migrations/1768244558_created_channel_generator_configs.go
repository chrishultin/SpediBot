package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"createRule": null,
			"deleteRule": null,
			"fields": [
				{
					"autogeneratePattern": "[a-z0-9]{15}",
					"hidden": false,
					"id": "text3208210256",
					"max": 15,
					"min": 15,
					"name": "id",
					"pattern": "^[a-z0-9]+$",
					"presentable": false,
					"primaryKey": true,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text2777438519",
					"max": 0,
					"min": 0,
					"name": "serverID",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text2766911014",
					"max": 0,
					"min": 0,
					"name": "channelID",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"hidden": false,
					"id": "select2634598704",
					"maxSelect": 1,
					"name": "nameFormat",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "select",
					"values": [
						"index",
						"owner"
					]
				},
				{
					"hidden": false,
					"id": "bool3020914760",
					"name": "enableRename",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "bool"
				},
				{
					"hidden": false,
					"id": "autodate2990389176",
					"name": "created",
					"onCreate": true,
					"onUpdate": false,
					"presentable": false,
					"system": false,
					"type": "autodate"
				},
				{
					"hidden": false,
					"id": "autodate3332085495",
					"name": "updated",
					"onCreate": true,
					"onUpdate": true,
					"presentable": false,
					"system": false,
					"type": "autodate"
				}
			],
			"id": "pbc_62587524",
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `IDX_UNIQ_CHAN_GEN_CONFIG` + "`" + ` ON ` + "`" + `channel_generator_configs` + "`" + ` (\n  ` + "`" + `serverID` + "`" + `,\n  ` + "`" + `channelID` + "`" + `\n)"
			],
			"listRule": null,
			"name": "channel_generator_configs",
			"system": false,
			"type": "base",
			"updateRule": null,
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_62587524")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
