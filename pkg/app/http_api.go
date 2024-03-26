package app

import (
	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/config"
	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/controller"
	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/db"
	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/db/migration"
	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/logger"
	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/repository"
	"github.com/wizeline/CA-Microservices-Go/internal/service"
)

type ApiHTTP struct {
	cfg    config.Config
	logger logger.ZeroLog
}

func NewApiHTTP(cfg config.Config, l logger.ZeroLog) (ApiHTTP, func(), error) {

	// Initialize database connection
	conn, err := db.NewPgConn(cfg.Database.Postgres)
	if err != nil {
		return ApiHTTP{}, nil, err
	}
	l.Log().Debug().Msg("database connection ready")

	// Run Migrations
	err = migration.Run(conn.DB(), []migration.Migration{
		migration.CreateUsersTable,
	}, l)
	if err != nil {
		return ApiHTTP{}, nil, err
	}

	// Clean Architecture
	userRepo := repository.NewUserRepoPg(conn.DB())
	userSvc := service.NewUserService(userRepo, l)
	_ = controller.NewUserController(userSvc)

	// TODO: imeplement the rest of application setup and start
	// - Router
	// - HttpServer

	return ApiHTTP{
			cfg:    cfg,
			logger: l,
		}, func() {
			if err := conn.Close(); err != nil {
				l.Log().Err(err).Msg("failed closing database connection")
			}
		}, nil

}

func (h ApiHTTP) Start() {
	h.logger.Log().Warn().Msg("http api start not implemented")
}
