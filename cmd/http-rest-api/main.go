package main

import (
	"os"

	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/config"
	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/logger"
	"github.com/wizeline/CA-Microservices-Go/pkg/app"
)

func main() {
	l := logger.NewZeroLog()
	cfg := config.NewConfig()

	api, err := app.NewApiHTTP(cfg, l)
	if err != nil {
		l.Log().Err(err).Msg("http rest api startup failed")
		os.Exit(1)
	}
	defer api.Shutdown()

	api.Start()
}
