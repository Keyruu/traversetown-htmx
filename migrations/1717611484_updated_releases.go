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

		// update
		edit_songtitle := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ukgsmdml",
			"name": "songtitle",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_songtitle); err != nil {
			return err
		}
		collection.Schema.AddField(edit_songtitle)

		// update
		edit_slug := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "7nzedc0l",
			"name": "slug",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_slug); err != nil {
			return err
		}
		collection.Schema.AddField(edit_slug)

		// update
		edit_artists := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ozvzeoqj",
			"name": "artists",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_artists); err != nil {
			return err
		}
		collection.Schema.AddField(edit_artists)

		// update
		edit_primaryColor := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "po7uzyfq",
			"name": "primaryColor",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": 4,
				"max": 7,
				"pattern": ""
			}
		}`), edit_primaryColor); err != nil {
			return err
		}
		collection.Schema.AddField(edit_primaryColor)

		// update
		edit_releaseDate := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ouqhoqpw",
			"name": "releaseDate",
			"type": "date",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_releaseDate); err != nil {
			return err
		}
		collection.Schema.AddField(edit_releaseDate)

		// update
		edit_cover := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xbkt3qb1",
			"name": "cover",
			"type": "file",
			"required": true,
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

		// update
		edit_apple := &schema.SchemaField{}
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
		}`), edit_apple); err != nil {
			return err
		}
		collection.Schema.AddField(edit_apple)

		// update
		edit_spotify := &schema.SchemaField{}
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
		}`), edit_spotify); err != nil {
			return err
		}
		collection.Schema.AddField(edit_spotify)

		// update
		edit_youtube := &schema.SchemaField{}
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
		}`), edit_youtube); err != nil {
			return err
		}
		collection.Schema.AddField(edit_youtube)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("3z7j67nhv66ad77")
		if err != nil {
			return err
		}

		// update
		edit_songtitle := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ukgsmdml",
			"name": "songtitle",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_songtitle); err != nil {
			return err
		}
		collection.Schema.AddField(edit_songtitle)

		// update
		edit_slug := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "7nzedc0l",
			"name": "slug",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_slug); err != nil {
			return err
		}
		collection.Schema.AddField(edit_slug)

		// update
		edit_artists := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ozvzeoqj",
			"name": "artists",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_artists); err != nil {
			return err
		}
		collection.Schema.AddField(edit_artists)

		// update
		edit_primaryColor := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "po7uzyfq",
			"name": "primaryColor",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": 4,
				"max": 7,
				"pattern": ""
			}
		}`), edit_primaryColor); err != nil {
			return err
		}
		collection.Schema.AddField(edit_primaryColor)

		// update
		edit_releaseDate := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ouqhoqpw",
			"name": "releaseDate",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_releaseDate); err != nil {
			return err
		}
		collection.Schema.AddField(edit_releaseDate)

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

		// update
		edit_apple := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qahhzy1q",
			"name": "apple",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_apple); err != nil {
			return err
		}
		collection.Schema.AddField(edit_apple)

		// update
		edit_spotify := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "smsvtvxr",
			"name": "spotify",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_spotify); err != nil {
			return err
		}
		collection.Schema.AddField(edit_spotify)

		// update
		edit_youtube := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "wjd1wuyp",
			"name": "youtube",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_youtube); err != nil {
			return err
		}
		collection.Schema.AddField(edit_youtube)

		return dao.SaveCollection(collection)
	})
}
