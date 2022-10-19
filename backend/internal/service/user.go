package service

import (
	"main/internal/models"
	"main/internal/repository"
)

type User interface {
	GetUserById(userId int) (models.User, error)
	GetUsers(username string) ([]models.User, error)
}

type UserService struct {
	repos repository.User
}

func newUserService(r repository.User) *UserService {
	return &UserService{
		repos: r,
	}
}

func (s *UserService) GetUserById(userId int) (models.User, error) {
	return s.repos.GetUserById(userId)
}

func (s *UserService) GetUsers(username string) ([]models.User, error) {
	return s.repos.GetUsers(username)
}
