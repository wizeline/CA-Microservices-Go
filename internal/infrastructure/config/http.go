package config

import (
	"fmt"
	"time"
)

// HTTPServer holds the config properties for the http server instance
type HTTPServer struct {
	host            string
	port            int
	shutdownTimeout time.Duration
}

// Address returns the TCP address for the server to listen on, in the form of "host:port"
func (h HTTPServer) Address() string {
	return fmt.Sprintf("%v:%d", h.host, h.port)
}

// ShutdownTimeout returns the timeout duration for shutting the server down
func (h HTTPServer) ShutdownTimeout() time.Duration {
	return h.shutdownTimeout
}
