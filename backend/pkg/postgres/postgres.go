package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

type dbConfig struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
	SSLMode  string
}

func InitDBConfig(host, port, user, dbName, password, sslMode string) *dbConfig {
	log.Println("init db config")
	return &dbConfig{
		Host:     host,
		Port:     port,
		User:     user,
		DbName:   dbName,
		Password: password,
		SSLMode:  sslMode,
	}
}

func InitDB(cfg *dbConfig) (*sql.DB, error) {
	log.Println("init db")

	uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DbName, cfg.Password, cfg.SSLMode)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
