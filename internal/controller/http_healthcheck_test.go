package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck_heartbeat(t *testing.T) {
	ctrl := NewHealthCheckHTTP()

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Get("/healthz", ctrl.heartbeat)
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"message\":\"OK\"}\n", rec.Body.String())
}
