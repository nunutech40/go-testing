package user_service

import (
	"go-testing/models/user"
	user_repository "go-testing/repository/user"
)

type UserService struct {
	repo *user_repository.UserRepository
}

func NewUserRepository(repo *user_repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]user.User, error) {
	return s.repo.GetUsers()
}
