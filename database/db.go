package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "myuser:mypassword@tcp(localhost:8000)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}

	// Check connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
