package service

import (
	"github.com/ramadani/adminarea"
	"github.com/ramadani/adminarea-cli/src"
	"github.com/ramadani/adminarea-cli/src/repository"
)

// CountryServiceImpl of CountryService contract
type CountryServiceImpl struct {
	repo      repository.Repository
	adminArea *adminarea.AdminArea
}

// Save country
func (s *CountryServiceImpl) Save() error {
	data := s.adminArea.GetCountry()

	_, err := s.repo.Save(&src.AdminArea{
		ID:   data.ID,
		Name: data.Name,
		Type: "country",
	})

	return err
}

// NewCountryService new country service
func NewCountryService(
	repo repository.Repository,
	adminArea *adminarea.AdminArea,
) *CountryServiceImpl {
	return &CountryServiceImpl{repo, adminArea}
}
