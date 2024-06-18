package utils

type ArtistInfo struct {
	Artist struct {
		Name  string `json:"name"`
		Stats struct {
			Listeners     string `json:"listeners"`
			Playcount     string `json:"playcount"`
			UserPlaycount string `json:"userplaycount"`
		}
	} `json:"artist"`
}
