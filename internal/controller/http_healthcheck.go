package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// We ensure the HTTP interface signature is satisfied by the HealthCheckHTTP implementation
var _ HTTP = &HealthCheckHTTP{}

// HealthCheckHTTP represents the system monitoring tool.
type HealthCheckHTTP struct{}

// NewHealthCheckHTTP returns a new HealthCheckHTTP implementation.
func NewHealthCheckHTTP() HealthCheckHTTP {
	return HealthCheckHTTP{}
}

// SetRoutes sets a fresh middleware stack to configure the handle functions of HealthCheckHTTP and mounts them to the given subrouter.
func (h HealthCheckHTTP) SetRoutes(r chi.Router) {
	r.Get("/healthz", h.heartbeat)
}

// heartbeat godoc
// @Summary Check if node is alive
// @Description  Check if node is alive
// @Tags         admin
// @Accept       json
// @Produce      json
// @Success      200  {object}  basicMessage
// @Failure      503  {object}  errHTTP
// @Router       /healthz [get]
func (h HealthCheckHTTP) heartbeat(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, basicMessage{
		Message: http.StatusText(http.StatusOK),
	})
}
