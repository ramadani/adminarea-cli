package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadani/adminarea-cli/command"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

type dbConfig struct {
	Driver string `yaml:"driver"`
	Dsn    string `yaml:"dsn"`
}

type config struct {
	DB dbConfig `yaml:"db"`
}

func (cog *config) readFromFile(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(file, cog)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cog := new(config)
	cog.readFromFile("adminarea_config.yml")

	db, err := sql.Open(cog.DB.Driver, cog.DB.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	setupCommand := command.NewSetupCommand(db)

	app := cli.NewApp()
	app.Name = "adminarea-cli"
	app.Usage = "Administrative Area Command Line Tool"
	app.Commands = []cli.Command{
		setupCommand,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
