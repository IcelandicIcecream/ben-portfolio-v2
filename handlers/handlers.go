package handlers

import (
	"github.com/a-h/templ"
	home "github.com/icelandicicecream/ben-portfolio-v2/views/pages/home"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return renderView(c, home.Home())
}

func renderView(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
