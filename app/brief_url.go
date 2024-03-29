package app

import (
	"log"
	"net"
	"strconv"
	"sync"

	"github.com/efimovalex/brief_url/adaptor/db"
)

// Config holds config values that will be read in via the Envy package.
// These are environment variables that will be modified by Chef.
type Config struct {
	Interface     string `name:"INTERFACE" example:"0.0.0.0"`
	Port          int    `name:"PORT" example:"50000"`
	JWTSigningKey string `name:"JWT_SIGNING_KEY" example:"23ASDcsSAFaFGjiGjAF2io3j"`
}

// Service contains private members to prepare this application/service.
type Service struct {
	*REST
}

// Start loads configs and starts listeners
func Start(config *Config, logger *log.Logger) error {
	// this service
	httpAddr := net.JoinHostPort(config.Interface, strconv.Itoa(config.Port))

	dbAdaptor, err := db.New()
	if err != nil {
		return err
	}

	MyApp := &Service{
		REST: &REST{
			Addr: httpAddr,
			Router: Routes(
				// add each dependent service as a dependency to the router
				dependencies{
					logger:    logger,
					config:    config,
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
