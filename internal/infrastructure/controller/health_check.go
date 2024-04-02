package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var _ HTTP = &HealthCheck{}

// HealthCheck is the system monitoring tool.
type HealthCheck struct{}

// NewHealthCheck returns a new HealthCheck implementation.
func NewHealthCheck() HealthCheck {
	return HealthCheck{}
}

// SetRoutes sets a fresh middleware stack for the HealthCheck handle functions and mounts them to the provided sub router.
func (h HealthCheck) SetRoutes(r chi.Router) {
	r.Get("/healthz", h.heartbeat)
}

// heartbeat is a handler function that checks the heartbeat of the API.
func (h HealthCheck) heartbeat(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, basicMessage{
		Message: http.StatusText(http.StatusOK),
	})
}
