package server

import (
	"github.com/labstack/echo/v4"
)

type Server struct{}

func (s Server) Start() (echoInstance *echo.Echo) {
	app := echo.New()
	s.HandleRoutes(app)

	return app
}
