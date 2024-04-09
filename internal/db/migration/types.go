package migration

import "database/sql"

// TODO: Implement migration functions for creating and dropping tables, modifying schema, or any other changes needed.
// Each migration function should be idempotent, meaning it can be run multiple times without causing issues.

var CreateUsersTable = Migration{
	name:     "CreateUsersTable",
	filename: "001_create_users_table.sql",
	Up: func(db *sql.DB, sqlContent string) error {
		_, err := db.Exec(sqlContent)
		return err
	},
	Down: func(db *sql.DB, sqlContent string) error {
		// Implement the rollback logic if needed
		_, err := db.Exec("DROP TABLE IF EXISTS users;")
		return err
	},
}
