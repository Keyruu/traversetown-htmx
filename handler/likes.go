package handler

import (
	"github.com/keyruu/traversetown-htmx/models"
	"github.com/keyruu/traversetown-htmx/utils"
	"github.com/keyruu/traversetown-htmx/views/likes"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
)

func (h *HandleController) LikesHandler(c echo.Context) error {
	fullstacks := []models.Fullstack{}
	err := h.getFullstacks("devops", &fullstacks)
	if err != nil {
		h.logger.Error("Something went wrong with getFullstacks", "error", err)
		return c.String(500, "error")
	}
	return utils.Render(c, 200, likes.Page(fullstacks))
}

func (h *HandleController) FullstackHandler(c echo.Context) error {
	fullstacks := []models.Fullstack{}
	stackType := c.QueryParam("type")
	err := h.getFullstacks(stackType, &fullstacks)
	if err != nil {
		return err
	}

	return utils.Render(c, 200, likes.Fullstack(fullstacks, stackType))
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
