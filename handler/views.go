package handler

import (
	"fmt"
	"net/http"

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

func (h *HandleController) getReleases(releases *[]models.Releases) error {
	err := h.dao.DB().
		NewQuery("SELECT * FROM releases ORDER BY `releaseDate` DESC").
		All(&releases)

	return err
}

func (h *HandleController) SidebarHandler(c echo.Context) error {
	return utils.Render(c, 200, components.Sidebar())
}

func (h *HandleController) IndexHandler(c echo.Context) error {
	releases := []models.Releases{}
	err := h.getReleases(&releases)
	if err != nil {
		return c.String(500, fmt.Sprintf("error: %e", err))
	}

	fullstacks := []models.Fullstack{}
	err = h.getFullstacks("devops", &fullstacks)
	if err != nil {
		return c.String(500, fmt.Sprintf("error: %e", err))
	}

	return utils.Render(c, 200, index.Page(releases, fullstacks))
}

func (h *HandleController) MusicRedirectHandler(c echo.Context) error {
	releases := []models.Releases{}
	err := h.getReleases(&releases)
	if err != nil {
		return c.String(500, fmt.Sprintf("error: %e", err))
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/music/"+releases[0].Slug)
}

func (h *HandleController) MusicHandler(c echo.Context) error {
	slug := c.PathParam("slug")
	if slug == "" {
		return c.String(404, "Not Found")
	}

	releases := []models.Releases{}
	err := h.getReleases(&releases)
	if err != nil {
		return c.String(500, fmt.Sprintf("error: %e", err))
	}

	selectedIndex := -1
	for i, release := range releases {
		if release.Slug == slug {
			selectedIndex = i
		}
	}

	if selectedIndex == -1 {
		return c.String(404, "Not Found")
	}

	return utils.Render(c, 200, music.Page(releases, selectedIndex))
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
