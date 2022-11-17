package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Auth
	User
	Product
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:    newAuthRepos(db),
		User:    newUserRepos(db),
		Product: newProductRepos(db),
		Admin:   newAdminRepository(db),
	}
}
