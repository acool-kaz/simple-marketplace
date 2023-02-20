package service

import (
	"context"
	"fmt"

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

func (ps *ProductService) Create(ctx context.Context, user models.User, product models.ProductCreate) (uint, error) {
	if user.Role == models.UserRoleInfo {
		product.UserId = user.Id
	}

	id, err := ps.productRepos.Create(ctx, product)
	if err != nil {
		return 0, fmt.Errorf("product service: create: %w", err)
	}

	return id, nil
}

func (ps *ProductService) GetAll(ctx context.Context) ([]models.Product, error) {
	products, err := ps.productRepos.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("product service: get all: %w", err)
	}

	return products, nil
}

func (ps *ProductService) GetOneBy(ctx context.Context) (models.Product, error) {
	product, err := ps.productRepos.GetOneBy(ctx)
	if err != nil {
		return models.Product{}, fmt.Errorf("product service: get one by: %w", err)
	}

	return product, nil
}

func (ps *ProductService) Update(ctx context.Context, user models.User, productId uint, product models.ProductUpdate) (models.Product, error) {
	if user.Role == models.UserRoleInfo {
		product, err := ps.productRepos.GetOneBy(context.WithValue(ctx, models.ProductId, productId))
		if err != nil {
			return models.Product{}, fmt.Errorf("product service: update: %w", err)
		}

		if product.UserId != productId {
			return models.Product{}, fmt.Errorf("product service: update: %w", models.ErrInvalidProduct)
		}
	}

	err := ps.productRepos.Update(ctx, productId, product)
	if err != nil {
		return models.Product{}, fmt.Errorf("product service: update: %w", err)
	}

	newProduct, err := ps.productRepos.GetOneBy(context.WithValue(ctx, models.ProductId, productId))
	if err != nil {
		return models.Product{}, fmt.Errorf("product service: update: %w", err)
	}

	return newProduct, nil
}

func (ps *ProductService) Delete(ctx context.Context, user models.User, productId uint) error {
	if user.Role == models.UserRoleInfo {
		product, err := ps.productRepos.GetOneBy(context.WithValue(ctx, models.ProductId, productId))
		if err != nil {
			return fmt.Errorf("product service: update: %w", err)
		}

		if product.UserId != productId {
			return fmt.Errorf("product service: update: %w", models.ErrInvalidProduct)
		}
	}

	err := ps.productRepos.Delete(ctx, productId)
	if err != nil {
		return fmt.Errorf("product service: delete: %w", err)
	}

	return nil
}
