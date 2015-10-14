package main

import (
	"log"

	"github.com/sendgrid/brief_url/app"
	"github.com/sendgrid/go-envy"
	"github.com/sendgrid/ln"
)

func main() {
	config := &app.Config{
		Interface: "0.0.0.0",
		Port:      "50111",
		MongoURLs: "http://localhost:8082",
	}

	ln.Info("starting with config", ln.Map{"config": config})

	if err := app.Start(config); err != nil {
		log.Fatal("Error occurred during startup:", err)
	}

	ln.Info("exiting.", nil)
}
