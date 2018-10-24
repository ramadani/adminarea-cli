package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ramadani/adminarea"

	"github.com/ramadani/adminarea-cli/src"
	"github.com/ramadani/adminarea-cli/src/repository"
	"github.com/urfave/cli"
)

// NewCountryCommand country command
func NewCountryCommand(repo repository.Repository) cli.Command {
	return cli.Command{
		Name:  "country",
		Usage: "Save a country",
		Action: func(c *cli.Context) error {
			cID := c.Args().First()
			if cID == "" {
				return errors.New("Please provide the country code id")
			}

			id := strings.ToUpper(cID)
			country := adminarea.New(id).GetCountry()
			_, err := repo.Save(&src.AdminArea{
				ID:   country.ID,
				Name: country.Name,
				Type: "COUNTRY",
			})

			if err != nil {
				return err
			}

			fmt.Println("Country has been saved")

			return nil
		},
	}
}
