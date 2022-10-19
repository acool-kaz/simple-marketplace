package repository

import (
	"fmt"
	"main/internal/models"

	"github.com/jmoiron/sqlx"
)

type User interface {
	GetUserById(userId int) (models.User, error)
	GetUsers(username string) ([]models.User, error)
}

type UserRepos struct {
	db *sqlx.DB
}

func newUserRepos(db *sqlx.DB) *UserRepos {
	return &UserRepos{
		db: db,
	}
}

func (r *UserRepos) GetUserById(userId int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = %d", users_table, userId)
	if err := r.db.Get(&user, query); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepos) GetUsers(username string) ([]models.User, error) {
	var users []models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = '%s' OR username = '%s'", users_table, username, username)
	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}
	return users, nil
}
