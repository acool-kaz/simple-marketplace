package service

import (
	"main/internal/models"
	"main/internal/repository"
)

type User interface {
	GetUserById(userId int) (models.User, error)
	GetUsers(username string) ([]models.User, error)
	GetAll() ([]models.User, error)
	DeleteUser(userId int) error
	UpdateUser(userId int, user models.User) error
}

type UserService struct {
	repos repository.User
}

func newUserService(r repository.User) *UserService {
	return &UserService{
		repos: r,
	}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repos.GetAll()
}

func (s *UserService) GetUserById(userId int) (models.User, error) {
	return s.repos.GetUserById(userId)
}

func (s *UserService) GetUsers(username string) ([]models.User, error) {
	return s.repos.GetUsers(username)
}

func (s *UserService) DeleteUser(userId int) error {
	return s.repos.DeleteUser(userId)
}

func (s *UserService) UpdateUser(userId int, user models.User) error {
	return s.repos.UpdateUser(userId, user)
}
