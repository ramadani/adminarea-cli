package repository

import "github.com/ramadani/adminarea-cli/src"

// Repository contract
type Repository interface {
	Save(adminArea *src.AdminArea) (string, error)
}
