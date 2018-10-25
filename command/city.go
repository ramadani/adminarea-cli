package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ramadani/adminarea"
	"github.com/ramadani/adminarea-cli/src"
	"github.com/ramadani/adminarea-cli/src/repository"
	"github.com/urfave/cli"
	pb "gopkg.in/cheggaaa/pb.v1"
)

// NewCityCommand new city command
func NewCityCommand(repo repository.Repository) cli.Command {
	return cli.Command{
		Name:  "city",
		Usage: "Save the cities/regencies of a country",
		Action: func(c *cli.Context) error {
			cID := c.Args().First()
			if cID == "" {
				return errors.New("Please provide the country code id")
			}

			id := strings.ToUpper(cID)
			aa := adminarea.New(id)
			cities, err := aa.GetRegencies()
			if err != nil {
				return err
			}

			count := len(cities)
			bar := pb.StartNew(count)

			for _, city := range cities {
				_, err = repo.Save(&src.AdminArea{
					ID:       city.ID,
					Name:     city.Name,
					Type:     "CITY",
					ParentID: city.ParentID,
				})

				if err != nil {
					return err
				}

				bar.Increment()
			}

			bar.FinishPrint(fmt.Sprintf("The cities of %s has been saved", aa.GetCountry().Name))

			return nil
		},
	}
}
