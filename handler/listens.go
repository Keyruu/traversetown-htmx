package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/keyruu/traversetown-htmx/utils"
	"github.com/keyruu/traversetown-htmx/views/listens"
	"github.com/labstack/echo/v5"
)

func (h *HandleController) ListensHandler(c echo.Context) error {
	return utils.Render(c, 200, listens.Page())
}

func (h *HandleController) SpotifyActivityHandler(c echo.Context) error {
	return utils.Render(c, 200, listens.SpotifyActivity())
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

	return utils.Render(c, 200, listens.LastfmAnswer(playcount, comment))
}
