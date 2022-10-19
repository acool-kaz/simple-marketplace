package service

import "main/internal/repository"

type Service struct {
	Auth
	User
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:    newAuthService(repos.Auth),
		User:    newUserService(repos.User),
		Product: newProductService(repos.Product),
	}
}
