package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

const defaultAppName = "camgo"

// Config is the representation of the application's configuration.
type Config struct {
	Application Application
	Server      HTTPServer
	Database    Database
}

func setDefaultConfig() {
	// Application configurations
	viper.SetDefault("application.name", defaultAppName)
	viper.SetDefault("application.version", "v0.0.0")
	// HTTP configurations
	viper.SetDefault("http.server.host", "localhost")
	viper.SetDefault("http.server.port", 8080)
	viper.SetDefault("http.server.shutdown.timeout", time.Second*15)
	// Database configurations
	viper.SetDefault("database.driver", "postgres")
	viper.SetDefault("database.postgres.host", "localhost")
	viper.SetDefault("database.postgres.port", 5432)
	viper.SetDefault("database.postgres.user", "postgres")
	viper.SetDefault("database.postgres.passwd", "")
	viper.SetDefault("database.postgres.dbname", defaultAppName)
}

// NewConfig creates a new Config instance
func NewConfig() Config {
	setDefaultConfig()
	viper.SetEnvPrefix(viper.GetString("application.name"))
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return Config{
		Application: Application{
			name:    viper.GetString("application.name"),
			version: viper.GetString("application.version"),
		},
		Server: HTTPServer{
			host:            viper.GetString("http.server.host"),
			port:            viper.GetInt("http.server.port"),
			shutdownTimeout: viper.GetDuration("http.server.shutdown.timeout"),
		},
		Database: Database{
			driver: viper.GetString("database.driver"),
			Postgres: PostgreSQL{
				Host:   viper.GetString("database.postgres.host"),
				Port:   viper.GetInt("database.postgres.port"),
				User:   viper.GetString("database.postgres.user"),
				DBName: viper.GetString("database.postgres.dbname"),
			},
		},
	}
}
