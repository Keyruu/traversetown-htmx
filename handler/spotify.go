package handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/cenkalti/dominantcolor"
	"github.com/keyruu/traversetown-htmx/config"
	"github.com/keyruu/traversetown-htmx/models"
	"github.com/keyruu/traversetown-htmx/utils"
	"github.com/keyruu/traversetown-htmx/views/listens"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/r3labs/sse/v2"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

type SpotifyController struct {
	server *sse.Server
	dao    *daos.Dao
	env    *config.Env
}

func NewSpotifyController(server *sse.Server, dao *daos.Dao, env *config.Env) *SpotifyController {
	return &SpotifyController{server: server, dao: dao, env: env}
}

type SpotifyToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func getTokenFromRefresh(refreshToken, clientId, clientSecret string) (*SpotifyToken, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	req, err := http.NewRequest(
		"POST",
		spotifyauth.TokenURL,
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	base64Client := base64.StdEncoding.EncodeToString([]byte(clientId + ":" + clientSecret))
	req.Header.Set("Authorization", "Basic "+base64Client)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response SpotifyToken

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *SpotifyController) SpotifyActivityTicker() {
	ctx := context.Background()

	expiresAt, client := c.refreshClient(ctx)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Printf("Token expires at: %v, now its %v\n", expiresAt, time.Now())
		log.Printf("Is it after expires? %s\n", strconv.FormatBool(time.Now().After(expiresAt)))
		if time.Now().After(expiresAt) {
			log.Printf("Refreshing token")
			expiresAt, client = c.refreshClient(ctx)
		}

		currentlyPlaying, err := client.PlayerCurrentlyPlaying(ctx)
		if err != nil {
			log.Printf("couldn't get features playlists: %v", err)
			continue
		}

		activity := &models.SpotifyActivity{}

		if currentlyPlaying.Item != nil && len(currentlyPlaying.Item.Album.Images) > 0 {
			activity = c.saveCurrent(currentlyPlaying)
		} else {
			err = c.dao.DB().NewQuery("SELECT * FROM spotify_activity").One(&activity)
			if err != nil {
				log.Printf("couldn't get spotify activity: %v", err)
			}

			activity.IsPlaying = false
		}

		log.Printf("Now playing: %s - %s\n", activity.TrackName, activity.ArtistName)
		var buffer bytes.Buffer
		listens.CurrentTrack(activity).Render(ctx, &buffer)
		event := &sse.Event{
			Data: buffer.Bytes(),
		}
		c.server.Publish("spotify", event)
	}
}

func (c *SpotifyController) refreshClient(ctx context.Context) (time.Time, *spotify.Client) {
	token, err := getTokenFromRefresh(c.env.SpotifyRefreshToken, c.env.SpotifyClientId, c.env.SpotifyClientSecret)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	expiresAt := time.Now().Add(time.Duration(token.ExpiresIn)*time.Second - 1*time.Minute)

	httpClient := spotifyauth.New().Client(ctx, &oauth2.Token{AccessToken: token.AccessToken})
	client := spotify.New(httpClient)
	return expiresAt, client
}

func (c *SpotifyController) saveCurrent(current *spotify.CurrentlyPlaying) *models.SpotifyActivity {
	spotifyActivity := models.SpotifyActivity{}

	err := c.dao.DB().NewQuery("SELECT * FROM spotify_activity").One(&spotifyActivity)
	if err != nil {
		log.Printf("couldn't get spotify activity: %v", err)
	}

	if spotifyActivity.SpotifyId != string(current.Item.ID) {
		resp, err := http.Get(utils.ResizeUrlJpeg(current.Item.Album.Images[0].URL, 100, 100))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		img, _, err := image.Decode(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		color := dominantcolor.Find(img)
		spotifyActivity.DominantColor = fmt.Sprintf("#%02x%02x%02x", color.R, color.G, color.B)
		log.Printf("Dominant color: %s", spotifyActivity.DominantColor)
	}

	spotifyActivity.SetCurrent(current)

	err = c.dao.Save(&spotifyActivity)
	if err != nil {
		log.Printf("couldn't save spotify activity: %v", err)
	}

	return &spotifyActivity
}
