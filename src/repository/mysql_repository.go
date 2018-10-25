package repository

import (
	"database/sql"
	"time"

	"github.com/ramadani/adminarea-cli/src"
)

// MySQLRepository contains deps
type MySQLRepository struct {
	db *sql.DB
}

// Save to mysql db
func (r *MySQLRepository) Save(adminArea *src.AdminArea) (string, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO administrative_areas (id, name, type, parent_id, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY
		UPDATE name = ?, type = ?, parent_id = ?, updated_at = ?
	`)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(adminArea.ID, adminArea.Name, adminArea.Type, adminArea.ParentID, now, now,
		adminArea.Name, adminArea.Type, adminArea.ParentID, now)

	if err != nil {
		return "", err
	}

	return adminArea.ID, nil
}

// NewMySQLRepository instance of mysql repository
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db}
}
