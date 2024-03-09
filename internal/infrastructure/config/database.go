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

// PostgreSQL holds and retrieves the config properties of the postgresql database instances.
type PostgreSQL struct {
	host   string
	port   int
	user   string
	passwd string
	dbname string
}

// Host returns the host value set for the postgresql instance.
func (pg PostgreSQL) Host() string {
	return pg.host
}

// Port returns the port value set for the postgresql instance.
func (pg PostgreSQL) Port() int {
	return pg.port
}

// User returns the username value set for the postgresql instance.
func (pg PostgreSQL) User() string {
	return pg.user
}

// Passwd returns the password value set for the postgresql instance.
func (pg PostgreSQL) Passwd() string {
	return pg.passwd
}

// DBName returns the database name value set for the postgresql instance.
func (pg PostgreSQL) DBName() string {
	return pg.dbname
}
