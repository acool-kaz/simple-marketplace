package service

import (
	"fmt"
	"main/internal/models"
	"main/internal/repository"
	"time"
)

type Product interface {
	GetAll() ([]models.Product, error)
	Create(product models.Product, userId int) error
	Find(find string) ([]models.Product, error)
	Delete(productId int) error
	Update(productId int, product models.Product) error
}

type ProductService struct {
	repos repository.Product
}

func newProductService(r repository.Product) *ProductService {
	return &ProductService{
		repos: r,
	}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repos.GetAll()
}

func (s *ProductService) Create(product models.Product, userId int) error {
	product.UserId = userId
	product.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	if product.Price <= 0 {
		return fmt.Errorf("%w: invalid price", models.ErrCreateProduct)
	}
	if len(product.Tags) == 0 {
		return fmt.Errorf("%w: one tag required", models.ErrCreateProduct)
	}
	return s.repos.Create(product)
}

func (s *ProductService) Delete(productId int) error {
	return s.repos.Delete(productId)
}

func (s *ProductService) Update(productId int, product models.Product) error {
	return s.repos.Update(productId, product)
}

func (s *ProductService) Find(find string) ([]models.Product, error) {
	return s.repos.Find(find)
}
