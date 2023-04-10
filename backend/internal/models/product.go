package models

import "mime/multipart"

type Product struct {
	Id               uint    `json:"id,omitempty"`
	UserId           uint    `json:"user_id,omitempty"`
	Name             string  `json:"name,omitempty"`
	ShortDescription string  `json:"short_description,omitempty"`
	Description      string  `json:"description,omitempty"`
	Tag              string  `json:"tag,omitempty"`
	Price            float64 `json:"price,omitempty"`
}

type ProductInfo struct {
	ProductId               uint     `json:"product_id,omitempty"`
	UserId                  uint     `json:"user_id,omitempty"`
	UserUsername            string   `json:"user_username,omitempty"`
	UserPhoneNumber         string   `json:"user_phone_number,omitempty"`
	ProductName             string   `json:"product_name,omitempty"`
	ProductShortDescription string   `json:"product_short_description,omitempty"`
	ProductDescription      string   `json:"product_description,omitempty"`
	ProductTag              string   `json:"product_tag,omitempty"`
	ProductPrice            float64  `json:"product_price,omitempty"`
	ProductPromoImage       string   `json:"product_promo_image,omitempty"`
	ProductImages           []string `json:"product_images,omitempty"`
}

type ProductCreate struct {
	UserId           uint                    `form:"user_id" json:"user_id,omitempty"`
	Name             string                  `form:"name" json:"name,omitempty" binding:"required"`
	ShortDescription string                  `form:"short_description" json:"short_description,omitempty" binding:"required"`
	Description      string                  `form:"description" json:"description,omitempty" binding:"required"`
	Tag              string                  `form:"tag" json:"tag,omitempty" binding:"required"`
	Price            float64                 `form:"price" json:"price,omitempty" binding:"required"`
	Images           []*multipart.FileHeader `form:"images"`
}

type ProductUpdate struct {
	Name             string  `json:"name,omitempty"`
	ShortDescription string  `json:"short_description,omitempty"`
	Description      string  `json:"description,omitempty"`
	Tag              string  `json:"tag,omitempty"`
	Price            float64 `json:"price,omitempty"`
}

var ProductSortParams = []string{"id", "user_id", "name", "description", "price"}

type ProductCtx string

const (
	ProductId               ProductCtx = "product_id"
	ProductUserId           ProductCtx = "product_user_id"
	ProductName             ProductCtx = "product_name"
	ProductShortDescription ProductCtx = "product_short_description"
	ProductDescription      ProductCtx = "product_description"
	ProductTag              ProductCtx = "product_tag"
	ProductPrice            ProductCtx = "product_price"

	ProductSortBy   ProductCtx = "product_sort_by"
	ProductFilterBy ProductCtx = "product_filter_by"
	ProductSearchBy ProductCtx = "product_search_by"
)
