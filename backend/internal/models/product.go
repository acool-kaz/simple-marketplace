package models

import "mime/multipart"

type Product struct {
	Id          uint    `json:"id,omitempty"`
	UserId      uint    `json:"user_id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
}

type ProductInfo struct {
	ProductId          uint    `json:"product_id,omitempty"`
	UserId             uint    `json:"user_id,omitempty"`
	UserUsername       string  `json:"user_username,omitempty"`
	UserPhoneNumber    string  `json:"user_phone_number,omitempty"`
	ProductName        string  `json:"product_name,omitempty"`
	ProductDescription string  `json:"product_description,omitempty"`
	ProductPrice       float64 `json:"product_price,omitempty"`
}

type ProductCreate struct {
	UserId      uint                    `from:"user_id" json:"user_id,omitempty"`
	Name        string                  `from:"name" json:"name,omitempty" binding:"required"`
	Description string                  `from:"description" json:"description,omitempty" binding:"required"`
	Price       float64                 `from:"price" json:"price,omitempty" binding:"required"`
	Images      []*multipart.FileHeader `form:"images"`
}

type ProductUpdate struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
}

var ProductSortParams = []string{"id", "user_id", "name", "description", "price"}

type ProductCtx string

const (
	ProductId          ProductCtx = "product_id"
	ProductUserId      ProductCtx = "product_user_id"
	ProductName        ProductCtx = "product_name"
	ProductDescription ProductCtx = "product_description"
	ProductPrice       ProductCtx = "product_price"

	ProductSortBy   ProductCtx = "product_sort_by"
	ProductFilterBy ProductCtx = "product_filter_by"
)
