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
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("3z7j67nhv66ad77")
		if err != nil {
			return err
		}

		// add
		new_coverHash := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "k8anr30c",
			"name": "coverHash",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_coverHash); err != nil {
			return err
		}
		collection.Schema.AddField(new_coverHash)

		// update
		edit_cover := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xbkt3qb1",
			"name": "cover",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [
					"image/webp",
					"image/png",
					"image/jpeg"
				],
				"thumbs": [
					"300x300",
					"600x600",
					"1000x1000",
					"1500x1500"
				],
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": false
			}
		}`), edit_cover); err != nil {
			return err
		}
		collection.Schema.AddField(edit_cover)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("3z7j67nhv66ad77")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("k8anr30c")

		// update
		edit_cover := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xbkt3qb1",
			"name": "cover",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [],
				"thumbs": [],
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": false
			}
		}`), edit_cover); err != nil {
			return err
		}
		collection.Schema.AddField(edit_cover)

		return dao.SaveCollection(collection)
	})
}
