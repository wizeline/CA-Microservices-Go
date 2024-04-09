package router

import (
	"reflect"

	"github.com/wizeline/CA-Microservices-Go/internal/config"
	"github.com/wizeline/CA-Microservices-Go/internal/controller"
	"github.com/wizeline/CA-Microservices-Go/internal/logger"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

// Chi configures a chi.Mux instance.
type Chi struct {
	basePath    string
	router      *chi.Mux
	controllers []controller.HTTP
	logger      logger.ZeroLog
}

// NewChi returns a Chi implementation.
// It allocates a pre-configured chi.Mux instance.
func NewChi(cfg config.Application, l logger.ZeroLog) Chi {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	return Chi{
		basePath:    cfg.BasePath(),
		router:      r,
		controllers: make([]controller.HTTP, 0),
		logger:      l,
	}
}

// Add appends the controller.HTTP given to the controller map list.
func (m *Chi) Add(ctrls ...controller.HTTP) {
	m.controllers = append(m.controllers, ctrls...)
}

// RegisterRoutes register the routes defined on the controller map list.
// All controller's routes are prefixed with the configured base path. e.g. "/api/v0"
func (m *Chi) RegisterRoutes() {
	if len(m.controllers) == 0 {
		m.logger.Log().Error().
			Str("error", "no http controllers found").
			Msg("register controller failed")
		return
	}
	m.router.Route(m.basePath, func(r chi.Router) {
		for _, ctrl := range m.controllers {
			name := reflect.TypeOf(ctrl).String()
			if ctrl == nil {
				m.logger.Log().Error().
					Str("error", "http controller not implemented").
					Str("controller", name).
					Msg("register controller failed")
				continue
			}

			ctrl.SetRoutes(r)
			m.logger.Log().Debug().
				Str("controller", name).
				Msg("registered http controller")
		}
	})
}

// Router returns the configured chi.Mux instance.
func (m *Chi) Router() *chi.Mux {
	return m.router
}
