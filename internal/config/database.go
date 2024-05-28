package config

// Database holds the configurations of the supported databases
type Database struct {
	Driver        string
	MigrationsDir string
	Postgres      PostgreSQL
}

// PostgreSQL holds the configuration values of the postgresql database instances.
type PostgreSQL struct {
	Host   string
	Port   int
	User   string
	Passwd string
	DBName string
}
