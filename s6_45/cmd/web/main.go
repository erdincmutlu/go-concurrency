package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const webPort = "80"

func main() {
	// connect to the database
	initDB()
	db.Ping()
	
	// create sessions

	// create channels

	// create waitgroup

	// set up the application config

	// set up mail

	// listen for web connections
}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to database")
	}
}

func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Print("postgres not yet ready...\n")
		} else  {
			log.Printf("Connected to the database\n")
			return connection
		}

		if count > 10 {
			return nil
		}

		log.Printf("Backing off for one second")
		time.Sleep(time.Second)

		count++
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db..Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
