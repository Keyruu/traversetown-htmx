package models

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

// ensures that the Article struct satisfy the models.Model interface
var _ models.Model = (*Releases)(nil)

type Releases struct {
	models.BaseModel

	Songtitle    string         `db:"songtitle" json:"songtitle"`
	Slug         string         `db:"slug" json:"slug"`
	Artists      string         `db:"artists" json:"artists"`
	PrimaryColor string         `db:"primaryColor" json:"primaryColor"`
	ReleaseDate  types.DateTime `db:"releaseDate" json:"releaseDate"`
	Cover        string         `db:"cover" json:"cover"`
	Apple        string         `db:"apple" json:"apple"`
	Spotify      string         `db:"spotify" json:"spotify"`
	Youtube      string         `db:"youtube" json:"youtube"`
	CoverHash    string         `db:"cover_hash" json:"cover_hash"`
}

func (m *Releases) TableName() string {
	return "releases" // the name of your collection
}
