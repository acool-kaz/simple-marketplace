package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Auth
	User
	Product
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:    newAuthRepos(db),
		User:    newUserRepos(db),
		Product: newProductRepos(db),
	}
}
