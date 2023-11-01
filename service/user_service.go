package service

import (
	"go-testing/models"
	"go-testing/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserRepository(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetUsers()
}
