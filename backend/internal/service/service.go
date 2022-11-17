package service

import "main/internal/repository"

type Service struct {
	Auth
	User
	Product
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:    newAuthService(repos.Auth),
		User:    newUserService(repos.User),
		Product: newProductService(repos.Product),
		Admin:   newAdminService(repos.Admin),
	}
}
