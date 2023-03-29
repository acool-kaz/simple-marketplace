package models

type Image struct {
	Id        uint   `json:"id,omitempty"`
	ProductId uint   `json:"product_id,omitempty"`
	Url       string `json:"url,omitempty"`
}

type ImageCreate struct {
	ProductId uint   `json:"product_id,omitempty"`
	Url       string `json:"url,omitempty"`
}

type ImageCtx string

const (
	ImageProductId ImageCtx = "image_product_id"
)
