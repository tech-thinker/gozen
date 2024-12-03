package runner

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"{{.PackageName}}/app/initializer"
	"{{.PackageName}}/app/router"
	"{{.PackageName}}/config"
	"{{.PackageName}}/instance"
	"{{.PackageName}}/logger"
)

// API is the interface for rest api runner
type API interface {
	Go(ctx context.Context, wg *sync.WaitGroup)
}

type api struct {
	cfg      config.Configuration
	instance instance.Instance
}

func (runner *api) Go(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithCancel(ctx)

	svc := initializer.Init(runner.cfg, runner.instance)

	routerV1 := router.Init(svc)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", runner.cfg.APIPort()),
		Handler:      routerV1,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Channel to listen for interrupt signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, os.Kill)

	go func() {
		<-stopChan
		fmt.Println("\nShutting down server...")
		cancel()

		// Create a context with timeout for the shutdown process
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			fmt.Printf("Server forced to shutdown: %v\n", err)
		} else {
			fmt.Println("Server stopped gracefully.")
		}
	}()

	logger.Log.Infof("Starting Rest API server...")
	logger.Log.Infof("Listening on server: http://127.0.0.1:%s", runner.cfg.APIPort())
	logger.Log.Infof("Health check URL: http://127.0.0.1:%s/ping", runner.cfg.APIPort())

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// NewAPI returns an instance of the REST API runner
func NewAPI(config config.Configuration, instance instance.Instance) API {
	return &api{
		cfg:      config,
		instance: instance,
	}
}
