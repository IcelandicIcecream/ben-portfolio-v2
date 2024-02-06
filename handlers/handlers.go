package handlers

import (
	"github.com/a-h/templ"
	"github.com/icelandicicecream/ben-portfolio-v2/model"
	"github.com/icelandicicecream/ben-portfolio-v2/views/components"
	home "github.com/icelandicicecream/ben-portfolio-v2/views/pages/home"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	state model.State
}

func (h Handler) Home(c echo.Context) error {
	return renderView(c, home.Home(h.state))
}

func (h Handler) OpenMobileNavbar(c echo.Context) error {
	return renderView(c, components.MobileNavBar())
}

func renderView(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
