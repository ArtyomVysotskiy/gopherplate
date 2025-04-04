// Package app configures and runs application.
package app

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"gopherplate/config"
	v1 "gopherplate/internal/controller/http"
	"gopherplate/internal/repo/persistent"
	"gopherplate/internal/repo/webapi"
	"gopherplate/internal/usecase/translation"
	"gopherplate/pkg/httpserver"
	"gopherplate/pkg/logger"
	"gopherplate/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal("app - Run - postgres.New", zap.Error(err))
	}
	defer pg.Close()

	// Use case
	translationUseCase := translation.New(
		persistent.New(pg),
		webapi.New(),
	)

	// HTTP Server
	httpServer := httpserver.New(
		httpserver.Port(cfg.HTTP.Port),
		httpserver.Prefork(cfg.HTTP.UsePreforkMode),
	)
	v1.NewRouter(httpServer.App, cfg, l, translationUseCase)

	// Start server
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal received", zap.String("signal", s.String()))
	case err := <-httpServer.Notify():
		l.Error("app - Run - httpServer.Notify", zap.Error(err))
	}

	// Shutdown
	if err := httpServer.Shutdown(); err != nil {
		l.Error("app - Run - httpServer.Shutdown", zap.Error(err))
	}
}
