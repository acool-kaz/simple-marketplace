package repository

import (
	"fmt"
	"main/internal/models"

	"github.com/jmoiron/sqlx"
)

type Product interface {
	GetAll() ([]models.Product, error)
	Create(product models.Product) error
	Delete(productId int) error
	Update(productId int, product models.Product) error
	Find(find string) ([]models.Product, error)
}

type ProductRepos struct {
	db *sqlx.DB
}

func newProductRepos(db *sqlx.DB) *ProductRepos {
	return &ProductRepos{
		db: db,
	}
}

func (r *ProductRepos) GetAll() ([]models.Product, error) {
	var products []models.Product
	query := fmt.Sprintf("SELECT * FROM %s", products_table)
	if err := r.db.Select(&products, query); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepos) Create(product models.Product) error {
	query := fmt.Sprintf("INSERT INTO %s(user_id, name, description, tags, price, created_at) VALUES ($1, $2, $3, $4, $5, $6)", products_table)
	if _, err := r.db.Exec(query, product.UserId, product.Name, product.Description, product.Tags, product.Price, product.CreatedAt); err != nil {
		return err
	}
	return nil
}

func (r *ProductRepos) Delete(productId int) error {
	query := fmt.Sprintf("DELTE FROM %s WHERE id = %d", products_table, productId)
	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (r *ProductRepos) Update(productId int, product models.Product) error {
	query := fmt.Sprintf("UPDATE %s SET user_id = %d, name = '%v', description = '%v', price = %d", products_table, product.UserId, product.Name, product.Description, product.Price)
	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (r *ProductRepos) Find(find string) ([]models.Product, error) {
	var products []models.Product
	query := fmt.Sprintf("SELECT * FROM %s WHERE name LIKE '%%%v%%' OR description LIKE '%%%v%%'", products_table, find, find)
	if err := r.db.Select(&products, query); err != nil {
		return nil, err
	}
	return products, nil
}
