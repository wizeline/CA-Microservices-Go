package controller

import (
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// We guarantee that the HTTP interface signature are satisfied by the SwaggerHTTP implementation
var _ HTTP = &SwaggerHTTP{}

type SwaggerHTTP struct{}

// NewHealthCheckHTTP returns a new SwaggerHTTP implementation.
func NewSwaggerHTTP() SwaggerHTTP {
	return SwaggerHTTP{}
}

// SetRoutes sets a fresh middleware stack to configure the handle functions of SwaggerHTTP and mounts them to the given subrouter.
func (s SwaggerHTTP) SetRoutes(r chi.Router) {
	r.Get("/swagger/*", httpSwagger.WrapHandler)
}
