package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// We guarantee that the requirements of the HTTP controller are met.
var _ HTTP = &HealthCheckController{}

// HealthCheckController is the system monitoring tool.
type HealthCheckController struct{}

// NewHealthCheck returns a new HealthCheck implementation.
func NewHealthCheck() *HealthCheckController {
	return &HealthCheckController{}
}

// SetRoutes sets a fresh middleware stack to configure the handle functions of HealthCheckController and mounts them to the given subrouter.
func (h HealthCheckController) SetRoutes(r chi.Router) {
	r.Get("/healthz", h.heartbeat)
}

// heartbeat is a handler function that checks the heartbeat of the API.
func (h HealthCheckController) heartbeat(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, basicMessage{
		Message: http.StatusText(http.StatusOK),
	})
}
