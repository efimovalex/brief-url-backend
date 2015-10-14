// Copyright 2014. SendGrid.
// Package MY_APP does blah blah blah.

package app

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/efimovalex/brief_url/adaptor/db"
)

// Config holds config values that will be read in via the Envy package.
// These are environment variables that will be modified by Chef.
type Config struct {
	Interface string `name:"INTERFACE" example:"0.0.0.0"`
	Port      int    `name:"PORT" example:"50110"`

	MongoURLs string `name:"MONGO_URLS" example:"http://localhost:8082"`
}

// Service contains private members to prepare this application/service.
type Service struct {
	*REST
}

// Start loads configs and starts listeners
func Start(config *Config) error {
	// this service
	httpAddr := net.JoinHostPort(config.Interface, strconv.Itoa(config.Port))
	httpListener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		return err
	}

	dbAdaptor := db.New(config.MongoURLs)

	MyApp := &Service{
		REST: &REST{
			Listener: httpListener,
			Router: Routes(
				// add each dependent service as a dependency to the router
				dependencies{
					logger:    *log.New(os.Stderr, "bried_url ", log.LstdFlags),
					dbAdaptor: dbAdaptor,
				}),
		},
	}

	return MyApp.Start()
}

// Start runs the entire service
// This implementation has an issue: if either crash, the whole service does not turn off
func (a *Service) Start() error {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		a.REST.StartHTTP()
	}()

	wg.Wait()

	return nil
}
