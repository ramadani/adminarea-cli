package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadani/adminarea-cli/command"
	"github.com/ramadani/adminarea-cli/src/repository"
	"github.com/ramadani/adminarea-cli/src/service"

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
	defer db.Close()

	repo := repository.NewMySQLRepository(db)
	countryService := service.NewCountryService(repo)

	setupCommand := command.NewSetupCommand(db)
	countryCommand := command.NewCountryCommand(countryService)

	app := cli.NewApp()
	app.Name = "adminarea-cli"
	app.Usage = "Administrative Area Command Line Tool"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		setupCommand,
		countryCommand,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
