package config

import (
	"fmt"
	"time"
)

// HTTPServer holds the config properties for the http server instance
type HTTPServer struct {
	Host            string
	Port            int
	ShutdownTimeout time.Duration
}

// Address returns the TCP address for the server to listen on, in the form of "host:port"
func (h HTTPServer) Address() string {
	return fmt.Sprintf("%v:%d", h.Host, h.Port)
}
