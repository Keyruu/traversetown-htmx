package models

import (
	"github.com/pocketbase/pocketbase/models"
)

// ensures that the Article struct satisfy the models.Model interface
var _ models.Model = (*Fullstack)(nil)

type Fullstack struct {
	models.BaseModel

	Type        string `db:"type" json:"type"`
	Logo        string `db:"logo" json:"logo"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Link        string `db:"link" json:"link"`
	Order       int    `db:"order" json:"order"`
}

func (m *Fullstack) TableName() string {
	return "fullstack" // the name of your collection
}
