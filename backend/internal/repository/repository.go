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
	imageTable   = "images"
)

type User interface {
	Create(ctx context.Context, user models.UserSignUp) (uint, error)
	GetAll(ctx context.Context) ([]models.User, error)
	GetOneBy(ctx context.Context) (models.User, error)
	Update(ctx context.Context, userId uint, user models.UserUpdate) error
	Delete(ctx context.Context, userId uint) error
}

type Product interface {
	Create(ctx context.Context, product models.ProductCreate) (uint, error)
	GetAll(ctx context.Context) ([]models.Product, error)
	GetOneBy(ctx context.Context) (models.Product, error)
	Update(ctx context.Context, productId uint, product models.ProductUpdate) error
	Delete(ctx context.Context, productId uint) error
}

type Image interface {
	Create(ctx context.Context, image models.ImageCreate) error
	GetAll(ctx context.Context) ([]models.Image, error)
}

type Repository struct {
	User    User
	Product Product
	Image   Image
}

func InitRepository(db *sql.DB) *Repository {
	log.Println("init repository")
	return &Repository{
		User:    newUserRepos(db),
		Product: newProductRepos(db),
		Image:   newImageRepos(db),
	}
}
