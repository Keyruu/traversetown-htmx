package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
)

type contextKey string

var EnvContextKey contextKey = "env"
var PathContextKey contextKey = "path"

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return ctx.HTML(500, err.Error())
	}

	return ctx.HTML(statusCode, buf.String())
}
