package main

import (
	"log"
	"os"

	"github.com/efimovalex/brief_url/app"
)

func main() {
	config := &app.Config{
		Interface: "0.0.0.0",
		Port:      50000,
	}

	log := *log.New(os.Stderr, "bried_url ", log.LstdFlags)

	log.Printf("starting with config: %v", map[string]app.Config{"config": *config})

	if err := app.Start(config, &log); err != nil {
		log.Fatal("Error occurred during startup: ", err)
	}

	log.Println("exiting.", nil)
}
