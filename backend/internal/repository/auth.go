package repository

import (
	"fmt"
	"main/internal/models"

	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(user models.User) error
	GetUser(username, password string) (models.User, error)
}

type AuthRepos struct {
	db *sqlx.DB
}

func newAuthRepos(db *sqlx.DB) *AuthRepos {
	return &AuthRepos{
		db: db,
	}
}

func (r *AuthRepos) CreateUser(user models.User) error {
	query := fmt.Sprintf("INSERT INTO %s (name, username, city, street, card_nums, card_m_y, password) VALUES($1, $2, $3, $4, $5, $6, $7)", users_table)
	if _, err := r.db.Exec(query, user.Name, user.Username, user.City, user.Street, user.CardNums, user.CardMY, user.Password); err != nil {
		return err
	}
	return nil
}

func (r *AuthRepos) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username='%s' AND password='%s'", users_table, username, password)
	if err := r.db.Get(&user, query); err != nil {
		return models.User{}, err
	}
	return user, nil
}
