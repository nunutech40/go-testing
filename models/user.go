package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"nama"`
	Age  int    `json:"umur"`
}
