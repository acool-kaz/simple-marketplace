package repository

import (
	"context"
	"database/sql"
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
	panic("not implemented") // TODO: Implement
}

func (pr *ProductRepos) GetOneBy(ctx context.Context) (models.Product, error) {
	panic("not implemented") // TODO: Implement
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
