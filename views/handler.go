package views

import (
	"github.com/a-h/templ"
	"github.com/keyruu/traversetown-htmx/models"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/daos"
)

type HandleController struct {
	dao *daos.Dao
}

func NewController(dao *daos.Dao) *HandleController {
	return &HandleController{dao: dao}
}

func (h *HandleController) IndexHandler(c echo.Context) error {
	releases := []models.Releases{}
	err := h.dao.DB().NewQuery("SELECT * FROM releases").All(&releases)
	if err != nil {
		return err
	}

	return Render(c, 200, Index(releases))
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
