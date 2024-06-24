package utils

import (
	"context"
	"fmt"

	"github.com/keyruu/traversetown-htmx/config"
	"github.com/keyruu/traversetown-htmx/models"
)

func GetEnv(ctx context.Context) *config.Env {
	return ctx.Value(EnvContextKey).(*config.Env)
}

func ReleasesImagePath(release models.Releases) string {
	return fmt.Sprintf("/api/files/releases/%s/%s", release.Id, release.Cover)
}

type InfoBoxIcon struct {
	Src         string
	Alt         string
	Description string
	Link        string
}
