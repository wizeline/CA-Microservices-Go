package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DBPgConn handles the PostgreSQL database connection.
type DBPgConn struct {
	db *sql.DB
}

// NewDBPgConn creates a new database connector instance.
func NewDBPgConn(host, port, user, password, dbname string) (*DBPgConn, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &DBPgConn{db}, nil
}

// Close closes the database connection.
func (conn *DBPgConn) Close() {
	conn.db.Close()
}

// DB returns the underlying *sql.DB instance.
func (conn *DBPgConn) DB() *sql.DB {
	return conn.db
}
