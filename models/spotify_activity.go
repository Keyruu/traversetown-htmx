package models

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/zmb3/spotify/v2"
)

// ensures that the Article struct satisfy the models.Model interface
var _ models.Model = (*SpotifyActivity)(nil)

type SpotifyActivity struct {
	models.BaseModel

	SpotifyId     string `db:"spotifyId" json:"spotifyId"`
	TrackName     string `db:"trackName" json:"trackName"`
	ArtistName    string `db:"artistName" json:"artistName"`
	CoverUrl      string `db:"coverUrl" json:"coverUrl"`
	DominantColor string `db:"dominantColor" json:"dominantColor"`
	SongLink      string `db:"songLink" json:"songLink"`
	IsPlaying     bool   `db:"isPlaying" json:"isPlaying"`
	ProgressMs    int    `db:"progressMs" json:"progressMs"`
	DurationMs    int    `db:"durationMs" json:"durationMs"`
	IsTooDark     bool   `db:"isTooDark" json:"isTooDark"`
}

func (m *SpotifyActivity) TableName() string {
	return "spotify_activity" // the name of your collection
}

func (m *SpotifyActivity) SetCurrent(current *spotify.CurrentlyPlaying) {
	track := current.Item

	m.IsPlaying = current.Playing
	m.ProgressMs = int(current.Progress)
	m.DurationMs = int(track.Duration)
	m.SpotifyId = string(track.ID)
	m.ArtistName = getArtistsString(track.Artists)
	m.TrackName = track.Name
	m.CoverUrl = track.Album.Images[0].URL
	m.SongLink = track.ExternalURLs["spotify"]
}

func getArtistsString(artists []spotify.SimpleArtist) string {
	artistsString := ""
	for _, artist := range artists {
		artistsString += artist.Name + ", "
	}
	return artistsString[:len(artistsString)-2]
}
