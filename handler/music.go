package handler

import (
	"fmt"
	"net/http"

	"github.com/keyruu/traversetown-htmx/models"
	"github.com/keyruu/traversetown-htmx/utils"
	"github.com/keyruu/traversetown-htmx/views/music"
	"github.com/labstack/echo/v5"
)

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

func (h *HandleController) getReleases(releases *[]models.Releases) error {
	err := h.dao.DB().
		NewQuery("SELECT * FROM releases ORDER BY `releaseDate` DESC").
		All(&releases)

	return err
}

func (h *HandleController) ReleasesSiteMap(urls *utils.Urlset) error {
	releases := []models.Releases{}
	err := h.getReleases(&releases)
	if err != nil {
		return err
	}

	for _, release := range releases {
		urls.AddUrl(utils.Url{Loc: fmt.Sprintf("%s/music/%s", h.env.BaseUrl, release.Slug)})
	}

	return nil
}
