package main

import (
	"authentication/data"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	// start with _ to import the package and not use it
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Docker allows multiple containers to listen to the same port
const webPort = "80"

// number of retries to connect to the database
var counts int64
const dbRetries int64 = 10

type Config struct {
	DB *sql.DB
	Models data.Models

}

func main() {
	log.Println("Starting authtentication service")

	// Add database connection here
	conn := connectToDB()

	if conn == nil {
		log.Panic("Could not connect to the database")
	}


	// set up configuration
	app := Config{
		DB: conn,
		Models: data.New(conn),
	}

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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// try to ping the database and return the error if it fails
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	// return the database connection and nil error
	return db, nil
}

func connectToDB() *sql.DB {
	// set up the database connection
	dsn := os.Getenv("DSN")

	// infinite for loop, stay there until the database connection is successful
	for {
		// try to open the database connection
		connection, err := openDB(dsn)
		if err != nil {
			// if there is an error, log it and try again
			log.Println("Postgres not yet ready...", err)
			counts++
		} else {
			// if we get here, the database connection was successful
			log.Println("Successfully connected to Postgres!")
			return connection
		}
		// if we have tried to connect to the database 10 times, panic
		if counts > dbRetries {
			log.Panic("Could not connect to the database")
			return nil
		}
		// sleep for 2 seconds before trying again
		log.Println("Sleeping for 2 seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}