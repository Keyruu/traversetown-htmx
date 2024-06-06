package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("3z7j67nhv66ad77")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("qahhzy1q")

		// remove
		collection.Schema.RemoveField("smsvtvxr")

		// remove
		collection.Schema.RemoveField("wjd1wuyp")

		// add
		new_apple := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "b6yx4hit",
			"name": "apple",
			"type": "url",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"exceptDomains": [],
				"onlyDomains": [
					"music.apple.com"
				]
			}
		}`), new_apple); err != nil {
			return err
		}
		collection.Schema.AddField(new_apple)

		// add
		new_spotify := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "s3p57xca",
			"name": "spotify",
			"type": "url",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"exceptDomains": [],
				"onlyDomains": [
					"open.spotify.com"
				]
			}
		}`), new_spotify); err != nil {
			return err
		}
		collection.Schema.AddField(new_spotify)

		// add
		new_youtube := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "lx7hmjuu",
			"name": "youtube",
			"type": "url",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"exceptDomains": [],
				"onlyDomains": [
					"www.youtube.com"
				]
			}
		}`), new_youtube); err != nil {
			return err
		}
		collection.Schema.AddField(new_youtube)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("3z7j67nhv66ad77")
		if err != nil {
			return err
		}

		// add
		del_apple := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qahhzy1q",
			"name": "apple",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_apple); err != nil {
			return err
		}
		collection.Schema.AddField(del_apple)

		// add
		del_spotify := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "smsvtvxr",
			"name": "spotify",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_spotify); err != nil {
			return err
		}
		collection.Schema.AddField(del_spotify)

		// add
		del_youtube := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "wjd1wuyp",
			"name": "youtube",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_youtube); err != nil {
			return err
		}
		collection.Schema.AddField(del_youtube)

		// remove
		collection.Schema.RemoveField("b6yx4hit")

		// remove
		collection.Schema.RemoveField("s3p57xca")

		// remove
		collection.Schema.RemoveField("lx7hmjuu")

		return dao.SaveCollection(collection)
	})
}
