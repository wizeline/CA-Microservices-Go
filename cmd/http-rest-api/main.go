package main

import (
	"os"

	"github.com/wizeline/CA-Microservices-Go/internal/config"
	"github.com/wizeline/CA-Microservices-Go/internal/logger"
	"github.com/wizeline/CA-Microservices-Go/pkg/app"
)

//	@title			CAM-Go REST API
//	@description	Code Accelerator Microservices REST API based on Golang.
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	CAM-Go
//	@contact.email	camgo@wizeline.com
//
// @securityDefinitions.basic	BasicAuth
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
