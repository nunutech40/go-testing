package user_repository

import (
	"database/sql"
	"go-testing/models/user"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetUsers() ([]user.User, error) {
	var users []user.User

	result, err := repo.db.Query("SELECT id, name as nama, age as umur from users")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var user user.User
		if err := result.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}

	return users, nil
}
