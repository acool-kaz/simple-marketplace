package repository

import (
	"fmt"
	"main/internal/models"

	"github.com/jmoiron/sqlx"
)

type Admin interface {
	GetAdmin(username, password string) (models.Admin, error)
}

type AdminRepository struct {
	db *sqlx.DB
}

func newAdminRepository(db *sqlx.DB) *AdminRepository {
	return &AdminRepository{
		db: db,
	}
}

func (r *AdminRepository) GetAdmin(username, password string) (models.Admin, error) {
	var admin models.Admin
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = '%v' AND password = '%v';", admins_table, username, password)
	if err := r.db.Get(&admin, query); err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}
