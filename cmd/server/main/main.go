package main

import (
	"log"
	"net/http"
	"fmt"	
)

func main() {
	app := Application{}
	app.Init()	
	
	fmt.Printf("Running a web server at %s \n",app.Config.addr)

    log.Fatal(http.ListenAndServe(app.Config.addr,app.Routes()))
}