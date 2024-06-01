package utils

import (
	"context"

	"github.com/keyruu/traversetown-htmx/config"
)

func GetEnv(ctx context.Context) *config.Env {
	return ctx.Value(EnvContextKey).(*config.Env)
}
