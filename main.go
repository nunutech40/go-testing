package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB
var err error

func main() {

	http.HandleFunc("/users", getUsers)
	http.ListenAndServe(":8080", nil)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []User

	result, err := db.Query("SELECT id, name, age from users")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}