package api

import (
	"github.com/wizeline/CA-Microservices-Go/internal/config"
	"github.com/wizeline/CA-Microservices-Go/internal/logger"
)

// SetSwaggerInfo overrides the default configuration of the swagger instance
func SetSwaggerInfo(cfg config.Application, l logger.ZeroLog) {
	SwaggerInfo.Version = cfg.Version()
	SwaggerInfo.Host = "" // force to retrieve the address (host:port) directly from the API
	SwaggerInfo.BasePath = cfg.BasePath()
	l.Log().Debug().Msg("configured swagger-info instance")
}
