package command

import (
	"fmt"

	"github.com/ramadani/adminarea-cli/src/migration"
	"github.com/urfave/cli"
)

// NewSetupCommand setup command
func NewSetupCommand(migrate migration.Migration) cli.Command {
	return cli.Command{
		Name:  "setup",
		Usage: "Create administrative areas table",
		Action: func(c *cli.Context) error {
			if err := migrate.Run(); err != nil {
				return err
			}

			fmt.Println("Setup done!")
			return nil
		},
	}
}
