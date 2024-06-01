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

		// remove
		collection.Schema.RemoveField("k8anr30c")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("3z7j67nhv66ad77")
		if err != nil {
			return err
		}

		// add
		del_coverHash := &schema.SchemaField{}
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
		}`), del_coverHash); err != nil {
			return err
		}
		collection.Schema.AddField(del_coverHash)

		return dao.SaveCollection(collection)
	})
}
