package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/acool-kaz/simple-marketplace/internal/repository"
	filesaver "github.com/acool-kaz/simple-marketplace/pkg/file_saver"
)

type ProductService struct {
	productRepos repository.Product
	imageRepos   repository.Image
	userRepos    repository.User

	imagePath string
}

func newProductService(productRepos repository.Product, imageRepos repository.Image, userRepos repository.User, imagePath string) *ProductService {
	return &ProductService{
		productRepos: productRepos,
		imageRepos:   imageRepos,
		userRepos:    userRepos,
		imagePath:    imagePath,
	}
}

func (ps *ProductService) Create(ctx context.Context, user models.User, product models.ProductCreate) (uint, error) {
	if user.Role == models.UserRoleInfo {
		product.UserId = user.Id
	}

	product.Name = strings.ToLower(product.Name)
	product.ShortDescription = strings.ToLower(product.ShortDescription)
	product.Description = strings.ToLower(product.Description)
	product.Tag = strings.ToLower(product.Tag)

	id, err := ps.productRepos.Create(ctx, product)
	if err != nil {
		return 0, fmt.Errorf("product service: create: %w", err)
	}

	for _, image := range product.Images {
		url, err := filesaver.SaveFile(ctx, "./static/product/", strconv.Itoa(int(id)), image)
		if err != nil {
			err = ps.Delete(ctx, user, id)
			if err != nil {
				log.Printf("product service: create: %v", err)
			}

			return 0, fmt.Errorf("product service: create: %w", err)
		}

		err = ps.imageRepos.Create(ctx, models.ImageCreate{ProductId: id, Url: ps.imagePath + url})
		if err != nil {
			err = ps.Delete(ctx, user, id)
			if err != nil {
				log.Printf("product service: create: %v", err)
			}

			return 0, fmt.Errorf("product service: create: %w", err)
		}
	}

	return id, nil
}

func (ps *ProductService) GetAllInfo(ctx context.Context) ([]models.ProductInfo, error) {
	products, err := ps.productRepos.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("product service: get all info: %w", err)
	}

	var productsInfo []models.ProductInfo

	for _, product := range products {
		user, err := ps.userRepos.GetOneBy(context.WithValue(ctx, models.UserId, product.UserId))
		if err != nil {
			return nil, fmt.Errorf("product service: get all info: %w", err)
		}

		productInfo := models.ProductInfo{
			ProductId:               product.Id,
			UserId:                  product.UserId,
			UserUsername:            user.Username,
			UserPhoneNumber:         user.PhoneNumber,
			ProductName:             product.Name,
			ProductShortDescription: product.ShortDescription,
			ProductDescription:      product.Description,
			ProductTag:              product.Tag,
			ProductPrice:            product.Price,
			ProductImages:           []string{},
		}

		images, err := ps.imageRepos.GetAll(context.WithValue(ctx, models.ImageProductId, product.Id))
		if err != nil {
			return nil, fmt.Errorf("product service: get all info: %w", err)
		}

		for _, img := range images {
			productInfo.ProductImages = append(productInfo.ProductImages, img.Url)
		}

		productInfo.ProductPromoImage = images[0].Url

		productsInfo = append(productsInfo, productInfo)
	}

	return productsInfo, nil
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

	err = os.RemoveAll(fmt.Sprintf("/static/product/%d", productId))
	if err != nil {
		log.Printf("product service: delete: %v", err)
	}

	return nil
}
