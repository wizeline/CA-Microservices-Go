package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/config"
)

// DBPgConn handles the PostgreSQL database connection.
type DBPgConn struct {
	db *sql.DB
}

// NewDBPgConn creates a new database connector instance.
func NewDBPgConn(cfg config.PostgreSQL) (*DBPgConn, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Passwd, cfg.DBName)

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
