package main

import (
	"log"
	"os"

	"github.com/icelandicicecream/ben-portfolio-v2/sentry"
	"github.com/icelandicicecream/ben-portfolio-v2/server"
)

func main() {
	sentryInstance := sentry.Sentry{}
	server := server.Server{}

	app := server.Start()
	err := sentryInstance.Init()
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentryInstance.Flush(2)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	err = app.Start(":" + port)
	if err != nil {
		sentryInstance.CaptureMessage("app.Start: " + err.Error())
		panic(err)
	}
}
