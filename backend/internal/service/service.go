package service

import (
	"context"
	"log"

	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/acool-kaz/simple-marketplace/internal/repository"
)

type User interface {
	Create(ctx context.Context, user models.UserSignUp) (uint, error)
	GetAll(ctx context.Context) ([]models.User, error)
	GetOneBy(ctx context.Context) (models.User, error)
	Update(ctx context.Context, userId uint, user models.UserUpdate) (models.User, error)
	Delete(ctx context.Context, userId uint) error
}

type Auth interface {
	SignUp(ctx context.Context, user models.UserSignUp) (uint, error)
	SignIn(ctx context.Context, user models.UserSignIn) (string, string, error)
	ParseToken(ctx context.Context, accessToken string) (uint, error)
}

type Service struct {
	Auth Auth
	User User
}

func InitService(repos *repository.Repository) *Service {
	log.Println("init service")
	return &Service{
		Auth: newAuthService(repos.User),
		User: newUserService(repos.User),
	}
}
