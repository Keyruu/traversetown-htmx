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

		collection, err := dao.FindCollectionByNameOrId("sdq5erepwu9t4ny")
		if err != nil {
			return err
		}

		// add
		new_isTooDark := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "tvzv59d1",
			"name": "isTooDark",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_isTooDark); err != nil {
			return err
		}
		collection.Schema.AddField(new_isTooDark)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("sdq5erepwu9t4ny")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("tvzv59d1")

		return dao.SaveCollection(collection)
	})
}
