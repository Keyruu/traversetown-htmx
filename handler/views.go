package handler

import (
	"log/slog"
	"net/url"

	"github.com/keyruu/traversetown-htmx/config"
	"github.com/keyruu/traversetown-htmx/models"
	"github.com/keyruu/traversetown-htmx/utils"
	"github.com/keyruu/traversetown-htmx/views/components"
	"github.com/keyruu/traversetown-htmx/views/index"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/daos"
)

type HandleController struct {
	dao    *daos.Dao
	env    *config.Env
	logger *slog.Logger
}

func NewController(dao *daos.Dao, env *config.Env, logger *slog.Logger) *HandleController {
	return &HandleController{dao: dao, env: env}
}

func (h *HandleController) SidebarHandler(c echo.Context) error {
	currentUrl := c.Request().Header.Get("Hx-Current-Url")
	url, err := url.Parse(currentUrl)
	path := ""
	if err == nil {
		path = url.Path
	}
	return utils.Render(c, 200, components.Sidebar(path))
}

func (h *HandleController) IndexHandler(c echo.Context) error {
	releases := []models.Releases{}
	err := h.getReleases(&releases)
	if err != nil {
		h.logger.Error("Something went wrong with getReleases", "error", err)
		return c.String(500, "error")
	}

	fullstacks := []models.Fullstack{}
	err = h.getFullstacks("devops", &fullstacks)
	if err != nil {
		h.logger.Error("Something went wrong with getFullstacks", "error", err)
		return c.String(500, "error")
	}

	return utils.Render(c, 200, index.Page(releases, fullstacks))
}
