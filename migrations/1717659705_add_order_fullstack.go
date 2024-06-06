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

		collection, err := dao.FindCollectionByNameOrId("yr60kima3ey0uk0")
		if err != nil {
			return err
		}

		// add
		new_order := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "h0kn2mig",
			"name": "order",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), new_order); err != nil {
			return err
		}
		collection.Schema.AddField(new_order)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("yr60kima3ey0uk0")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("h0kn2mig")

		return dao.SaveCollection(collection)
	})
}
