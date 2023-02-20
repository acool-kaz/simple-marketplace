package models

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
	UserId      uint    `json:"user_id,omitempty"`
	Name        string  `json:"name,omitempty" binding:"required"`
	Description string  `json:"description,omitempty" binding:"required"`
	Price       float64 `json:"price,omitempty" binding:"required"`
}

type ProductUpdate struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
}

type ProductCtx string

const (
	ProductId          ProductCtx = "product_id"
	ProductUserId      ProductCtx = "product_user_id"
	ProductName        ProductCtx = "product_name"
	ProductDescription ProductCtx = "product_description"
	ProductPrice       ProductCtx = "product_price"
)
