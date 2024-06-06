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
		new_link := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "cqmmkx8f",
			"name": "link",
			"type": "url",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"exceptDomains": [],
				"onlyDomains": []
			}
		}`), new_link); err != nil {
			return err
		}
		collection.Schema.AddField(new_link)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("yr60kima3ey0uk0")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("cqmmkx8f")

		return dao.SaveCollection(collection)
	})
}
