package main

import (
	"log"
	"os"
	
	"github.com/urfave/cli/v2"
	
	"40104/SecretsManager/cmd/cli_client/models"
)

func main() {
	app := &Application{CLI: &cli.App{}, Env: &models.Env{}, Param: &models.Params{}}
	app.EnvVariable("configs/app.env")
	app.CLI = app.Setup()

	if err := app.CLI.Run(os.Args); err != nil {
        log.Fatal(err)
	}
}
