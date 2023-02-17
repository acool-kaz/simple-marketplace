package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/acool-kaz/simple-marketplace/internal/models"
)

const (
	userTable    = "users"
	productTable = "products"
)

type User interface {
	Create(ctx context.Context, user models.UserSignUp) (uint, error)
	GetAll(ctx context.Context) ([]models.User, error)
	GetOneBy(ctx context.Context) (models.User, error)
	Update(ctx context.Context, userId uint, user models.UserUpdate) error
	Delete(ctx context.Context, userId uint) error
}

type Product interface {
	Create(ctx context.Context, userId uint, product models.ProductCreate) (uint, error)
	GetAll(ctx context.Context) ([]models.ProductInfo, error)
	GetOneBy(ctx context.Context) (models.ProductInfo, error)
	Update(ctx context.Context, productId uint, product models.ProductUpdate) error
	Delete(ctx context.Context, productId uint) error
}

type Repository struct {
	User    User
	Product Product
}

func InitRepository(db *sql.DB) *Repository {
	log.Println("init repository")
	return &Repository{
		User:    newUserRepos(db),
		Product: newProductRepos(db),
	}
}
