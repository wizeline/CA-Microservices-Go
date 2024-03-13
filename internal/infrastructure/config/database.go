package config

// Database holds the configurations of the supported databases
type Database struct {
	driver   string
	Postgres PostgreSQL
}

// Driver returns the configured database driver.
func (db Database) Driver() string {
	return db.driver
}

// PostgreSQL holds the configuration values of the postgresql database instances.
type PostgreSQL struct {
	Host   string
	Port   int
	User   string
	Passwd string
	DBName string
}
