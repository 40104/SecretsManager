package main

import (
	"log"
	"os"
	//"fmt"

	"github.com/joho/godotenv"
	
	"40104/SecretsManager/cmd/server/models"
	"40104/SecretsManager/cmd/server/controllers"
)

type Application struct {
	//Model *models.DBModel
	Controller *controllers.Controller
	Config *Config
}

type Config struct {
	connection_string string
	addr string
	key string

}

func (app *Application) Init() {
	
	if err := godotenv.Load("configs/app.env"); err != nil {
    	log.Fatal(err)
    }

	config := &Config{
		connection_string:	os.Getenv("CONNECTION_STRING"),
		addr:	os.Getenv("ADDR"),

	}
	app.Config = config

	
	app.Controller = &controllers.Controller{
		JWT_secret:	[]byte(os.Getenv("JWT")),
	}

	db := app.Controller.DBModel.ConnectDB(app.Config.connection_string)

	//defer db.Close()

	dbmodel := &models.DBModel{
		DB: db,
		Key: os.Getenv("KEY"),
	}
	
	app.Controller.DBModel = dbmodel
	

	app.Controller.DBModel.InitDB()
	
}

