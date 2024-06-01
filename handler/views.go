package handler

import (
	"github.com/keyruu/traversetown-htmx/models"
	"github.com/keyruu/traversetown-htmx/utils"
	"github.com/keyruu/traversetown-htmx/views/index"
	"github.com/keyruu/traversetown-htmx/views/music"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/daos"
)

type HandleController struct {
	dao *daos.Dao
}

func NewController(dao *daos.Dao) *HandleController {
	return &HandleController{dao: dao}
}

func (h *HandleController) IndexHandler(c echo.Context) error {
	releases := []models.Releases{}
	err := h.dao.DB().NewQuery("SELECT * FROM releases").All(&releases)
	if err != nil {
		return err
	}

	return utils.Render(c, 200, index.Page(releases))
}

func (h *HandleController) MusicHandler(c echo.Context) error {
	return utils.Render(c, 200, music.Page())
}
