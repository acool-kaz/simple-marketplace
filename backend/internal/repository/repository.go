package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/acool-kaz/simple-marketplace/internal/models"
)

const (
	userTable = "users"
)

type User interface {
	Create(ctx context.Context, user models.UserSignUp) (uint, error)
	GetAll(ctx context.Context) ([]models.User, error)
	GetOneBy(ctx context.Context) (models.User, error)
	Update(ctx context.Context, userId uint, user models.UserUpdate) (models.User, error)
	Delete(ctx context.Context, userId uint) error
}

type Repository struct {
	User User
}

func InitRepository(db *sql.DB) *Repository {
	log.Println("init repository")
	return &Repository{
		User: newUserRepos(db),
	}
}
