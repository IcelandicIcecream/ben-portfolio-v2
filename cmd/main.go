package main

import (
	"github.com/icelandicicecream/ben-portfolio-v2/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/styles", "styles")

	app.GET("/", handlers.Home)

	err := app.Start(":3000")
	if err != nil {
		panic(err)
	}
}
