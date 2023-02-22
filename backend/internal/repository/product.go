package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/acool-kaz/simple-marketplace/internal/models"
)

type ProductRepos struct {
	db *sql.DB
}

func newProductRepos(db *sql.DB) *ProductRepos {
	return &ProductRepos{
		db: db,
	}
}

func (pr *ProductRepos) Create(ctx context.Context, product models.ProductCreate) (uint, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s
			(user_id, name, description, price)
		VALUES
			('%d', '%s', '%s', '%f')
		RETURNING id;`,
		productTable,
		product.UserId, product.Name, product.Description, product.Price,
	)

	var id uint
	err := pr.db.QueryRowContext(ctx, query).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("product repos: create: %w", err)
	}

	return id, err
}

func (pr *ProductRepos) GetAll(ctx context.Context) ([]models.Product, error) {
	query := fmt.Sprintf(`
		SELECT
			*
		FROM %s;`,
		productTable,
	)

	row, err := pr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("product repos: get all: %w", err)
	}
	defer row.Close()

	var products []models.Product

	for row.Next() {
		var product models.Product

		err = row.Scan(&product.Id, &product.UserId, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, fmt.Errorf("product repos: get all: %w", err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (pr *ProductRepos) GetOneBy(ctx context.Context) (models.Product, error) {
	argsStr := []string{}

	ctxKeys := []interface{}{models.ProductId, models.ProductUserId, models.ProductName, models.ProductDescription, models.ProductPrice}

	for _, ctxKey := range ctxKeys {
		ctxValue := ctx.Value(ctxKey)
		if ctxValue != nil {
			ctxKeyString := string(ctxKey.(models.ProductCtx))
			ctxKeyString = strings.TrimPrefix(ctxKeyString, "product_")

			argsStr = append(argsStr, fmt.Sprintf("%s = '%v'", ctxKeyString, ctxValue))
		}
	}

	whereCondition := ""
	if len(argsStr) != 0 {
		whereCondition = "WHERE " + strings.Join(argsStr, " AND ")
	}

	query := fmt.Sprintf(`
		SELECT
			*
		FROM %s
		%s`,
		productTable,
		whereCondition,
	)

	row := pr.db.QueryRowContext(ctx, query)

	var product models.Product

	err := row.Scan(&product.Id, &product.UserId, &product.Name, &product.Description, &product.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = models.ErrProductNotFound
		}
		return models.Product{}, fmt.Errorf("product repos: get one by: %w", err)
	}

	return product, nil
}

func (pr *ProductRepos) Update(ctx context.Context, productId uint, product models.ProductUpdate) error {
	argsStr := []string{}

	if product.Name != "" {
		argsStr = append(argsStr, fmt.Sprintf("name = '%s'", product.Name))
	}

	if product.Description != "" {
		argsStr = append(argsStr, fmt.Sprintf("description = '%s'", product.Description))
	}

	if product.Price > 0 {
		argsStr = append(argsStr, fmt.Sprintf("price = '%f'", product.Price))
	}

	updateQuery := ""

	if len(argsStr) != 0 {
		updateQuery = strings.Join(argsStr, ", ")
	}

	query := fmt.Sprintf(`
		UPDATE
			%s
		SET
			%s
		WHERE id = '%d'
		`,
		userTable,
		updateQuery,
		productId,
	)

	_, err := pr.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("product repos: update: %w", err)
	}

	return nil
}

func (pr *ProductRepos) Delete(ctx context.Context, productId uint) error {
	query := fmt.Sprintf(`
		DELETE FROM
			%s
		WHERE id = '%d'`,
		productTable,
		productId,
	)

	_, err := pr.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("product repos: delete: %w", err)
	}

	return nil
}
