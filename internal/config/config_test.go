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
					Name:    defaultAppName,
					Version: "v0.0.0",
				},
				HTTPServer: HTTPServer{
					Host:            "localhost",
					Port:            8080,
					ShutdownTimeout: 15000000000,
				},
				Database: Database{
					Driver:        "postgres",
					MigrationsDir: "migrations/v1",
					Postgres: PostgreSQL{
						Host:   "pgdb",
						Port:   5432,
						User:   defaultAppName + "user",
						Passwd: defaultAppName + "p44s5W0rD",
						DBName: defaultAppName,
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
