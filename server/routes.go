package server

import (
	"github.com/icelandicicecream/ben-portfolio-v2/handlers"
	"github.com/labstack/echo/v4"
)

func (s Server) HandleRoutes(app *echo.Echo) {
	handlers := handlers.Handler{}
	app.Static("/styles", "styles")
	app.Static("/assets", "assets")
	app.GET("/", handlers.Home)
	app.GET("/check-health", handlers.CheckHealth)
	app.GET("/open-mobile-navbar", handlers.OpenMobileNavbar)
}
