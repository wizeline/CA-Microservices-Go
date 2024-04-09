package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/wizeline/CA-Microservices-Go/internal/config"
	"github.com/wizeline/CA-Microservices-Go/internal/controller"
	"github.com/wizeline/CA-Microservices-Go/internal/db"
	"github.com/wizeline/CA-Microservices-Go/internal/db/migration"
	"github.com/wizeline/CA-Microservices-Go/internal/logger"
	"github.com/wizeline/CA-Microservices-Go/internal/repository"
	"github.com/wizeline/CA-Microservices-Go/internal/router"
	"github.com/wizeline/CA-Microservices-Go/internal/service"
)

type ApiHTTP struct {
	cfg    config.HTTPServer
	dbConn *db.PgConn
	server *http.Server
	logger logger.ZeroLog
}

func NewApiHTTP(cfg config.Config, l logger.ZeroLog) (ApiHTTP, error) {

	// Initialize database connection
	dbConn, err := db.NewPgConn(cfg.Database.Postgres)
	if err != nil {
		return ApiHTTP{}, err
	}
	l.Log().Debug().Msg("database connection ready")

	// Run Migrations
	err = migration.Run(dbConn.DB(), []migration.Migration{
		migration.CreateUsersTable,
	}, l)
	if err != nil {
		return ApiHTTP{}, err
	}

	// User dependencies
	userRepo := repository.NewUserRepoPg(dbConn.DB())
	userSvc := service.NewUserService(userRepo, l)

	// Router
	r := router.NewChi(cfg.Application, l)
	r.Add(controller.NewHealthCheck(), controller.NewUserController(userSvc))
	r.RegisterRoutes()

	return ApiHTTP{
		cfg:    cfg.HTTPServer,
		dbConn: dbConn,
		server: &http.Server{
			Handler:      r.Router(),
			Addr:         cfg.HTTPServer.Address(),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
			IdleTimeout:  time.Second * 60,
		},
		logger: l,
	}, nil

}

// Start runs the http API and quits doing a grateful shutdown.
// To stop the server you must send a syscall.SIGINT signal usually through `CTRL+C`.
func (h ApiHTTP) Start() {
	go func() {
		h.logger.Log().Info().Msgf("running http server on %v", h.cfg.Address())
		err := h.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			h.logger.Log().Fatal().Err(err).Msg("http server startup failed")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

// Shutdown performs tasks of safely shutting down processes and closing connections.
func (h ApiHTTP) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), h.cfg.ShutdownTimeout())
	defer cancel()

	if err := h.server.Shutdown(ctx); err != nil {
		h.logger.Log().Error().Err(err).Msg("http server graceful shutdown failed")
	}

	if err := h.dbConn.Close(); err != nil {
		h.logger.Log().Err(err).Msg("failed closing database connection")
	}

	h.logger.Log().Info().Msg("http api shutdown gracefully")
	os.Exit(0)
}
