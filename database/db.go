package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "myuser:mypassword@tcp(localhost:8000)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}

	// Check connection
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
