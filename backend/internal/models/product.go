package models

import (
	"errors"

	"github.com/lib/pq"
)

type Product struct {
	Id          int            `json:"id" db:"id"`
	UserId      int            `json:"user_id" db:"user_id"`
	Name        string         `json:"name" db:"name"`
	Description string         `json:"description" db:"description"`
	Tags        pq.StringArray `json:"tags" db:"tags"`
	Price       int            `json:"price" db:"price"`
	CreatedAt   string         `json:"created_at" db:"created_at"`
}

var (
	ErrCreateProduct = errors.New("cant create product")
)
