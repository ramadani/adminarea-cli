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

// NewVillageCommand new village command
func NewVillageCommand(repo repository.Repository) cli.Command {
	return cli.Command{
		Name:  "village",
		Usage: "Save the villages of a country",
		Action: func(c *cli.Context) error {
			cID := c.Args().First()
			if cID == "" {
				return errors.New("Please provide the country code id")
			}

			id := strings.ToUpper(cID)
			aa := adminarea.New(id)
			villages, err := aa.GetVillages()
			if err != nil {
				return err
			}

			count := len(villages)
			bar := pb.StartNew(count)

			for _, village := range villages {
				_, err = repo.Save(&src.AdminArea{
					ID:       village.ID,
					Name:     village.Name,
					Type:     "VILLAGE",
					ParentID: village.ParentID,
				})

				if err != nil {
					return err
				}

				bar.Increment()
			}

			bar.FinishPrint(fmt.Sprintf("The villages of %s has been saved", aa.GetCountry().Name))

			return nil
		},
	}
}
