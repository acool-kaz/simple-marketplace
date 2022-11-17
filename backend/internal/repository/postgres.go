package repository

import (
	"fmt"
	"main/config"

	"github.com/jmoiron/sqlx"
)

const (
	users_table    = "users"
	admins_table   = "admins"
	products_table = "products"
)

func InitDB(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.DB_Host, cfg.DB_Port, cfg.DB_Username, cfg.DB_Name, cfg.DB_Password, cfg.DB_SSLMode))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
