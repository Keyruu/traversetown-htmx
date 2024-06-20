package handler

import (
	"github.com/keyruu/traversetown-htmx/utils"
	"github.com/keyruu/traversetown-htmx/views/imprint"
	"github.com/labstack/echo/v5"
)

func (h *HandleController) ImprintHandler(c echo.Context) error {
	return utils.Render(c, 200, imprint.Page())
}
