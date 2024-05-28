package migration

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wizeline/CA-Microservices-Go/internal/logger"
)

type Migration struct {
	name     string
	Filename string
	Up       func(db *sql.DB, sqlContent string) error
	Down     func(db *sql.DB, sqlContent string) error
}

// Run applies the given migration functions
func Run(db *sql.DB, migrationsDir string, migrations []Migration, l logger.ZeroLog) error {
	for _, m := range migrations {
		filePath := filepath.Join(migrationsDir, m.Filename)
		l.Log().Debug().Str("name", m.name).Str("file_path", filePath).
			Msg("applying migration")

		sqlContent, err := os.ReadFile(filePath)
		if err != nil {
			// TODO: convert it to error migration type
			return fmt.Errorf("failed reading SQL file %s: %s", m.Filename, err)
		}

		if err := m.Up(db, string(sqlContent)); err != nil {
			// TODO: convert it to error migration type
			return fmt.Errorf("failed applying migration %s: %s", m.name, err)
		}
	}
	return nil
}
