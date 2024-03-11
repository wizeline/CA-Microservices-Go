package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name    string
		envVars []string
		exp     Config
	}{
		// TODO: Add more test cases
		{
			name:    "Default",
			envVars: nil,
			exp: Config{
				Application: Application{
					name:    defaultAppName,
					version: "v0.0.0",
				},
				Server: HTTPServer{
					host:            "localhost",
					port:            8080,
					shutdownTimeout: 15000000000,
				},
				Database: Database{
					driver: "postgres",
					Postgres: PostgreSQL{
						host:   "localhost",
						port:   5432,
						user:   "postgres",
						passwd: "",
						dbname: defaultAppName,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := NewConfig()
			assert.Equal(t, tt.exp, out)
		})
	}
}
