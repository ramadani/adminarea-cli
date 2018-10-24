package migration

import (
	"database/sql"

	"github.com/ramadani/adminarea-cli/resources"
)

// MySQLMigration mysql migration
type MySQLMigration struct {
	db *sql.DB
}

// Run the migration
func (m *MySQLMigration) Run() error {
	query, err := resources.Asset("create_administrative_areas_table.sql")
	if err != nil {
		return err
	}

	if _, err := m.db.Exec(string(query[:])); err != nil {
		return err
	}

	return nil
}

// NewMySQLMigration new mysql migration
func NewMySQLMigration(db *sql.DB) *MySQLMigration {
	return &MySQLMigration{db}
}
