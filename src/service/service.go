package service

// CountryService contract
type CountryService interface {
	SaveByID(id string) error
}
