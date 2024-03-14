package migration

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/logger"
)

const migrationsDir = "internal/infrastructure/db/migration/v1"

type Migration struct {
	name     string
	filename string
	Up       func(db *sql.DB, sqlContent string) error
	Down     func(db *sql.DB, sqlContent string) error
}

// Run applies the given migration functions
func Run(db *sql.DB, migrations []Migration, l logger.ZeroLog) error {
	for _, m := range migrations {
		l.Log().Debug().Str("name", m.name).Msg("applying migration")

		path := filepath.Join(migrationsDir, m.filename)
		sqlContent, err := os.ReadFile(path)
		if err != nil {
			// TODO: convert it to error migration type
			return fmt.Errorf("failed reading SQL file %s: %s", m.filename, err)
		}

		if err := m.Up(db, string(sqlContent)); err != nil {
			// TODO: convert it to error migration type
			return fmt.Errorf("failed applying migration %s: %s", m.name, err)
		}
	}
	return nil
}
