package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "yr60kima3ey0uk0",
			"created": "2024-06-05 18:25:03.140Z",
			"updated": "2024-06-05 18:25:03.140Z",
			"name": "fullstack",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "drujowdw",
					"name": "type",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"devops",
							"frontend",
							"backend"
						]
					}
				},
				{
					"system": false,
					"id": "6akmh2u7",
					"name": "logo",
					"type": "file",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"mimeTypes": [
							"image/png",
							"image/jpeg",
							"image/webp"
						],
						"thumbs": [],
						"maxSelect": 1,
						"maxSize": 5242880,
						"protected": false
					}
				},
				{
					"system": false,
					"id": "zffxtllk",
					"name": "name",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "q69zj5o6",
					"name": "description",
					"type": "editor",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"convertUrls": false
					}
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("yr60kima3ey0uk0")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
