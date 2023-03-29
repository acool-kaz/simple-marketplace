package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/acool-kaz/simple-marketplace/internal/models"
)

type ImageRepos struct {
	db *sql.DB
}

func newImageRepos(db *sql.DB) *ImageRepos {
	return &ImageRepos{
		db: db,
	}
}

func (ir *ImageRepos) Create(ctx context.Context, image models.ImageCreate) error {
	query := fmt.Sprintf(`
		INSERT INTO %s
			(product_id, url)
		VALUES
			('%d', '%s')
		;`,
		imageTable,
		image.ProductId, image.Url,
	)

	_, err := ir.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("image repos: creat: %w", err)
	}

	return nil
}

func (ir *ImageRepos) GetAll(ctx context.Context) ([]models.Image, error) {
	argsStr := []string{}

	ctxKeys := []interface{}{models.ImageProductId}

	for _, ctxKey := range ctxKeys {
		ctxValue := ctx.Value(ctxKey)
		if ctxValue != nil {
			ctxKeyString := string(ctxKey.(models.ImageCtx))
			ctxKeyString = strings.TrimPrefix(ctxKeyString, "image_")

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
		%s;`,
		imageTable,
		whereCondition,
	)

	rows, err := ir.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("image repos: get all: %w", err)
	}
	defer rows.Close()

	var images []models.Image

	for rows.Next() {
		var image models.Image

		err = rows.Scan(&image.Id, &image.ProductId, &image.Url)
		if err != nil {
			return nil, fmt.Errorf("image repos: get all: %w", err)
		}

		images = append(images, image)
	}

	return images, nil
}
