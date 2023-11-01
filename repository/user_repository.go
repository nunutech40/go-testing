package repository

import (
	"database/sql"
	"go-testing/models"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User

	result, err := repo.db.Query("SELECT id, name as nama, age as umur from users")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var user models.User
		if err := result.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}

	return users, nil
}
