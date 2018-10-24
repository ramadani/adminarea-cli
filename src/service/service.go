package service

import "github.com/ramadani/adminarea-cli/src"

// CountryService contract
type CountryService interface {
	Save(data *src.AdminArea) error
}
