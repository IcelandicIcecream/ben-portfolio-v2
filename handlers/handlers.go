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
	return Render(c, home.Home(h.state))
}

func (h Handler) OpenMobileNavbar(c echo.Context) error {
	return Render(c, components.MobileNavBar())
}

func (h Handler) CheckHealth(c echo.Context) error {
	return c.String(200, "The echo server is up and running! 🚀")
}

func Render(c echo.Context, t templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(c.Request().Context(), c.Response().Writer)
}
