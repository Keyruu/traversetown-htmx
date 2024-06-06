package handler

import (
	"github.com/keyruu/traversetown-htmx/models"
	"github.com/keyruu/traversetown-htmx/utils"
	"github.com/keyruu/traversetown-htmx/views/components"
	"github.com/keyruu/traversetown-htmx/views/index"
	"github.com/keyruu/traversetown-htmx/views/music"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
)

type HandleController struct {
	dao *daos.Dao
}

func NewController(dao *daos.Dao) *HandleController {
	return &HandleController{dao: dao}
}

func (h *HandleController) getFullstacks(stackType string, fullstacks *[]models.Fullstack) error {
	err := h.dao.DB().
		NewQuery("SELECT * FROM fullstack WHERE `type` = {:type} ORDER BY `order`").
		Bind(dbx.Params{
			"type": stackType,
		}).
		All(&fullstacks)

	return err
}

func (h *HandleController) IndexHandler(c echo.Context) error {
	releases := []models.Releases{}
	err := h.dao.DB().NewQuery("SELECT * FROM releases ORDER BY `releaseDate`").All(&releases)
	if err != nil {
		return c.Render(500, "error", err)
	}

	fullstacks := []models.Fullstack{}
	err = h.getFullstacks("devops", &fullstacks)
	if err != nil {
		return c.Render(500, "error", err)
	}

	return utils.Render(c, 200, index.Page(releases, fullstacks))
}

func (h *HandleController) MusicHandler(c echo.Context) error {
	return utils.Render(c, 200, music.Page())
}

func (h *HandleController) FullstackHandler(c echo.Context) error {
	fullstacks := []models.Fullstack{}
	stackType := c.QueryParam("type")
	err := h.getFullstacks(stackType, &fullstacks)
	if err != nil {
		return err
	}

	return utils.Render(c, 200, components.Fullstack(fullstacks, stackType))
}
