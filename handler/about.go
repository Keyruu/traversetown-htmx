package handler

import (
	"github.com/keyruu/traversetown-htmx/utils"
	"github.com/keyruu/traversetown-htmx/views/about"
	"github.com/labstack/echo/v5"
)

func (h *HandleController) AboutHandler(c echo.Context) error {
	return utils.Render(c, 200, about.Page())
}
