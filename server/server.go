package server

import (
	"github.com/icelandicicecream/ben-portfolio-v2/logs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var log = logs.Logger()

type Server struct{}

func (s Server) Start() (echoInstance *echo.Echo) {
	app := echo.New()

	config := middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				log.Error().Msgf("%v | %s: '%s' -- %s", v.Status, v.Method, v.URI, v.Error.Error())
			}
			log.Info().Msgf("%v | %s: '%s'", v.Status, v.Method, v.URI)
			return nil
		},
	}
	app.Use(middleware.RequestLoggerWithConfig(config))
	s.HandleRoutes(app)

	return app
}
