package service

import (
	"context"
	"fmt"
	"log"

	"github.com/acool-kaz/simple-marketplace/internal/config"
	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/acool-kaz/simple-marketplace/internal/repository"
)

type Auth interface {
	SignUp(ctx context.Context, user models.UserSignUp) (uint, error)
	SignIn(ctx context.Context, user models.UserSignIn) (string, string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.Token, error)
}

type User interface {
	Create(ctx context.Context, user models.UserSignUp) (uint, error)
	GetAll(ctx context.Context) ([]models.User, error)
	GetOneBy(ctx context.Context) (models.User, error)
	Update(ctx context.Context, userId uint, user models.UserUpdate) (models.User, error)
	Delete(ctx context.Context, userId uint) error
}

type Product interface {
	GetAllInfo(ctx context.Context) ([]models.ProductInfo, error)
	Create(ctx context.Context, user models.User, product models.ProductCreate) (uint, error)
	GetAll(ctx context.Context) ([]models.Product, error)
	GetOneBy(ctx context.Context) (models.Product, error)
	Update(ctx context.Context, user models.User, productId uint, product models.ProductUpdate) (models.Product, error)
	Delete(ctx context.Context, user models.User, productId uint) error
}

type Service struct {
	Auth    Auth
	User    User
	Product Product
}

func InitService(repos *repository.Repository, cfg *config.Config) *Service {
	log.Println("init service")
	return &Service{
		Auth:    newAuthService(repos.User),
		User:    newUserService(repos.User),
		Product: newProductService(repos.Product, repos.Image, repos.User, fmt.Sprintf("%s:%s", cfg.HttpConfig.Host, cfg.HttpConfig.Port)),
	}
}
