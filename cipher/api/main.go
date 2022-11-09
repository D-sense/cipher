package main

import (
	"appital-service/api/handler"
	"appital-service/foundation/logger"
	"context"
	"errors"
	"fmt"
	"github.com/ardanlabs/conf/v3"
	"github.com/dimfeld/httptreemux/v5"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// build is the git version of this program. It is set using build flags in the makefile.
var build = "develop"

func main() {
	// Construct the application logger.
	log, err := logger.New("CIPHER-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	// =========================================================================
	// Configuration
	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:10s"`
			IdleTimeout     time.Duration `conf:"default:120s"`
			ShutdownTimeout time.Duration `conf:"default:20s"`
			APIHost         string        `conf:"default:0.0.0.0:8080"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "copyright information here",
		},
	}

	const prefix = "AGENT"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	// =========================================================================
	// Start API Service

	log.Infow("startup", "status", "initializing V1 Cipher API support")

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Construct the mux for the API calls.
	router := httptreemux.NewContextMux()

	group := router.NewGroup("/api")
	group.GET("/v1/encrypt", handler.Encrypt)
	group.GET("/v1/decrypt", handler.Decrypt)

	// Construct a server to service the requests against the mux.
	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      router,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog:     zap.NewStdLog(log.Desugar()),
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for api requests.
	go func() {
		log.Infow("startup", "status", "api router started", "host", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Infow("shutdown", "status", "shutdown started", "signal", sig)
		defer log.Infow("shutdown", "status", "shutdown complete", "signal", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		// Asking listener to shut down and shed load.
		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
