package command

import (
	"database/sql"
	"fmt"

	"github.com/ramadani/adminarea-cli/resources"
	"github.com/urfave/cli"
)

type migration struct {
	db *sql.DB
}

func (m *migration) Run() error {
	query, err := resources.Asset("create_administrative_areas_table.sql")
	if err != nil {
		return err
	}

	if _, err := m.db.Exec(string(query[:])); err != nil {
		return err
	}

	return nil
}

// NewSetupCommand setup command
func NewSetupCommand(db *sql.DB) cli.Command {
	return cli.Command{
		Name:  "setup",
		Usage: "Create administrative areas table",
		Action: func(c *cli.Context) error {
			migrate := migration{db}
			if err := migrate.Run(); err != nil {
				return err
			}

			fmt.Println("Setup done!")
			return nil
		},
	}
}
