package service

import (
	"context"

	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/acool-kaz/simple-marketplace/internal/repository"
)

type ProductService struct {
	productRepos repository.Product
}

func newProductService(productRepos repository.Product) *ProductService {
	return &ProductService{
		productRepos: productRepos,
	}
}

func (ps *ProductService) Create(ctx context.Context, product models.ProductCreate) (uint, error) {
	panic("not implemented") // TODO: Implement
}

func (ps *ProductService) GetAll(ctx context.Context) ([]models.Product, error) {
	panic("not implemented") // TODO: Implement
}

func (ps *ProductService) GetOneBy(ctx context.Context) (models.Product, error) {
	panic("not implemented") // TODO: Implement
}

func (ps *ProductService) Update(ctx context.Context, productId uint, product models.ProductUpdate) (models.Product, error) {
	panic("not implemented") // TODO: Implement
}

func (ps *ProductService) Delete(ctx context.Context, productId uint) error {
	panic("not implemented") // TODO: Implement
}
