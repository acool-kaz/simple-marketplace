package repository

import (
	"fmt"
	"main/internal/models"

	"github.com/jmoiron/sqlx"
)

type User interface {
	GetUserById(userId int) (models.User, error)
	GetUsers(username string) ([]models.User, error)
	GetAll() ([]models.User, error)
	DeleteUser(userId int) error
	UpdateUser(userId int, user models.User) error
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

func (r *UserRepos) GetAll() ([]models.User, error) {
	var users []models.User
	query := fmt.Sprintf("SELECT * FROM %s;", users_table)
	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepos) DeleteUser(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d", users_table, userId)
	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (r *UserRepos) UpdateUser(userId int, user models.User) error {
	query := fmt.Sprintf("UPDATE %s SET name='%v', username='%v', city='%v', street='%v', card_nums='%v', password='%v' WHERE id = %d", users_table, user.Name, user.Username, user.City, user.Street, user.CardNums, user.Password, userId)
	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}
