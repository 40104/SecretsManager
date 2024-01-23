package main
// Import packages
import (
	"log"
	"os"

	"github.com/joho/godotenv"
	
	"40104/SecretsManager/cmd/server/models"
	"40104/SecretsManager/cmd/server/controllers"
)
// Class Application
type Application struct {
	Controller *controllers.Controller
	Config *Config
}
// Subclass Config
type Config struct {
	connection_string string
	addr string
	key string

}
// Init function
func (app *Application) Init() {
	// Import config file
	if err := godotenv.Load("configs/app.env"); err != nil {
    	log.Fatal(err) // Check the error
    }
	// Init Config class
	config := &Config{
		connection_string:	os.Getenv("CONNECTION_STRING"), // Set connection string
		addr:	os.Getenv("ADDR"), // Set address

	}
	// Add Config to the Application class
	app.Config = config 
	// Init Controller
	app.Controller = &controllers.Controller{
		JWT_secret:	[]byte(os.Getenv("JWT")), // Set JWT key
	}
	// Init DB connection
	db := app.Controller.DBModel.ConnectDB(app.Config.connection_string)
	// Add DB model class
	dbmodel := &models.DBModel{
		DB: db, // Set Db connection
		Key: os.Getenv("KEY"), // Set encryption key
	}
	// Add Controller to the Application class
	app.Controller.DBModel = dbmodel
	// Init DB migration
	app.Controller.DBModel.InitDB()
}

