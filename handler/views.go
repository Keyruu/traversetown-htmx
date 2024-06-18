package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/keyruu/traversetown-htmx/config"
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
	env *config.Env
}

func NewController(dao *daos.Dao, env *config.Env) *HandleController {
	return &HandleController{dao: dao, env: env}
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

func (h *HandleController) getLastfmPlaycount(artist string) (int, error) {
	resp, err := http.Get(
		fmt.Sprintf("%s?method=artist.getinfo&artist=%s&api_key=%s&format=json&username=keyruu",
			h.env.LastfmUrl, url.QueryEscape(artist), h.env.LastfmApiKey))
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	artistInfo := utils.ArtistInfo{}
	err = json.NewDecoder(resp.Body).Decode(&artistInfo)
	if err != nil {
		return 0, err
	}

	playcount, err := strconv.Atoi(artistInfo.Artist.Stats.UserPlaycount)
	if err != nil {
		return 0, err
	}

	return playcount, nil
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

func (h *HandleController) LastfmHandler(c echo.Context) error {
	artist := c.FormValue("artist")
	if artist == "" {
		return c.String(400, "artist is required")
	}

	playcount, err := h.getLastfmPlaycount(artist)
	if err != nil {
		return c.String(500, fmt.Sprintf("error: %e", err))
	}

	lowerArtist := strings.ToLower(artist)
	comment := "Yeah..."
	if lowerArtist == "brakence" {
		comment =
			"This is my all-time favorite artist! brakence is a very big inspiration for me!"
	} else if lowerArtist == "keshi" {
		comment =
			"keshi is probably my second favorite artist! I love his music!"
	} else if playcount > 2000 {
		comment = "I love them! They are in my top 10!"
	} else if playcount > 1000 {
		comment =
			"Oh boy! This is an awesome artist! Probably in my top 20."
	} else if playcount > 500 {
		comment = "I like them! Hasn't made it to my top 10 though."
	} else if playcount > 100 {
		comment = "I do like them! But I do not listen to them that often."
	} else if playcount > 10 {
		comment =
			"I don't really know about them... Maybe I should listen to them more?"
	}

	return utils.Render(c, 200, components.LastfmAnswer(playcount, comment))
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
