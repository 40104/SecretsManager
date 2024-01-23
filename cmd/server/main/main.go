package main
// Import packages
import (
	"log"
	"net/http"
	"fmt"	
)

// Start program 
func main() {
	// Init Application class
	app := Application{}
	// Init base variables
	app.Init()	
	
	fmt.Printf("Running a web server at %s \n",app.Config.addr)
	// Start web server
    log.Fatal(http.ListenAndServe(app.Config.addr,app.Routes()))
}