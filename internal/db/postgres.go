package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/wizeline/CA-Microservices-Go/internal/config"
)

// PgConn handles the PostgreSQL database connection.
type PgConn struct {
	db *sql.DB
}

// NewPgConn creates a new database connector instance.
func NewPgConn(cfg config.PostgreSQL) (*PgConn, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host(), cfg.Port(), cfg.User(), cfg.Passwd(), cfg.DBName())

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

	return &PgConn{db}, nil
}

// Close closes the database connection.
func (conn *PgConn) Close() error {
	return conn.db.Close()
}

// DB returns the underlying *sql.DB instance.
func (conn *PgConn) DB() *sql.DB {
	return conn.db
}
