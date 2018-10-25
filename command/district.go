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

// NewDistrictCommand new city command
func NewDistrictCommand(repo repository.Repository) cli.Command {
	return cli.Command{
		Name:  "district",
		Usage: "Save the districts of a country",
		Action: func(c *cli.Context) error {
			cID := c.Args().First()
			if cID == "" {
				return errors.New("Please provide the country code id")
			}

			id := strings.ToUpper(cID)
			aa := adminarea.New(id)
			districts, err := aa.GetDistricts()
			if err != nil {
				return err
			}

			count := len(districts)
			bar := pb.StartNew(count)

			for _, district := range districts {
				_, err = repo.Save(&src.AdminArea{
					ID:       district.ID,
					Name:     district.Name,
					Type:     "DISTRICT",
					ParentID: district.ParentID,
				})

				if err != nil {
					return err
				}

				bar.Increment()
			}

			bar.FinishPrint(fmt.Sprintf("The districts of %s has been saved", aa.GetCountry().Name))

			return nil
		},
	}
}
