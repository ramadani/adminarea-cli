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

// NewProvinceCommand new province command
func NewProvinceCommand(repo repository.Repository) cli.Command {
	return cli.Command{
		Name:  "province",
		Usage: "Save the provinces of a country",
		Action: func(c *cli.Context) error {
			cID := c.Args().First()
			if cID == "" {
				return errors.New("Please provide the country code id")
			}

			id := strings.ToUpper(cID)
			aa := adminarea.New(id)
			provinces, err := aa.GetProvinces()
			if err != nil {
				return err
			}

			for _, province := range provinces {
				_, err = repo.Save(&src.AdminArea{
					ID:       province.ID,
					Name:     province.Name,
					Type:     "PROVINCE",
					ParentID: province.ParentID,
				})

				if err != nil {
					return err
				}
			}

			fmt.Println(fmt.Sprintf("The provinces of country %s has been saved", aa.GetCountry().Name))

			return nil
		},
	}
}
