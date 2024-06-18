package handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/keyruu/traversetown-htmx/config"
	"github.com/keyruu/traversetown-htmx/views/components"
	"github.com/r3labs/sse/v2"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

type SpotifyController struct {
	server *sse.Server
	env    *config.Env
}

func NewSpotifyController(server *sse.Server, env *config.Env) *SpotifyController {
	return &SpotifyController{server: server, env: env}
}

func getTokenFromRefresh(refreshToken, clientId, clientSecret string) (*oauth2.Token, error) {
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

	var response oauth2.Token

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *SpotifyController) SpotifyActivityTicker() {
	ctx := context.Background()

	token, err := getTokenFromRefresh(c.env.SpotifyRefreshToken, c.env.SpotifyClientId, c.env.SpotifyClientSecret)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		currentlyPlaying, err := client.PlayerCurrentlyPlaying(ctx)
		if err != nil {
			log.Fatalf("couldn't get features playlists: %v", err)
		}
		if currentlyPlaying.Item != nil {
			fmt.Printf("Now playing: %s\n", currentlyPlaying.Item.Name)
			var buffer bytes.Buffer
			components.CurrentTrack(*currentlyPlaying).Render(ctx, &buffer)
			event := &sse.Event{
				Data: buffer.Bytes(),
			}
			c.server.Publish("time", event)
		}
	}
}
