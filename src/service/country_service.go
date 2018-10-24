package service

import (
	"github.com/ramadani/adminarea-cli/src"
	"github.com/ramadani/adminarea-cli/src/repository"
)

// CountryServiceImpl of CountryService contract
type CountryServiceImpl struct {
	repo repository.Repository
}

// Save country
func (s *CountryServiceImpl) Save(data *src.AdminArea) error {
	_, err := s.repo.Save(data)

	return err
}

// NewCountryService new country service
func NewCountryService(repo repository.Repository) *CountryServiceImpl {
	return &CountryServiceImpl{repo}
}
