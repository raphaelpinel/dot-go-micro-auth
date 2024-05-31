package main

import (
	"authentication/data"
	"database/sql"
	"log"
	"net/http"
)

// Docker allows multiple containers to listen to the same port
const webPort = "80"

type Config struct {
	DB *sql.DB
	Models data.Models

}

func main() {
	log.Println("Starting authtentication service")

	// TODO: Add database connection here

	// set up configuration
	app := Config{}

	// web server
	srv := &http.Server{
		Addr: ":" + webPort,
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}