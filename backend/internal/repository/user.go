package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/acool-kaz/simple-marketplace/internal/models"
)

type UserRepos struct {
	db *sql.DB
}

func newUserRepos(db *sql.DB) *UserRepos {
	return &UserRepos{
		db: db,
	}
}

func (ur *UserRepos) Create(ctx context.Context, user models.UserSignUp) (uint, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s 
			(first_name, second_name, email, username, password) 
		VALUES 
			('%s', '%s', '%s', '%s', '%s') 
		RETURNING id;`,
		userTable,
		user.FirstName, user.SecondName, user.Email, user.Username, user.Password,
	)

	var id uint
	err := ur.db.QueryRowContext(ctx, query).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("user repos: create: %w", err)
	}

	return id, err
}

func (ur *UserRepos) GetAll(ctx context.Context) ([]models.User, error) {
	query := fmt.Sprintf(`
		SELECT
			*
		FROM %s`,
		userTable,
	)

	rows, err := ur.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("user repos: get all: %w", err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.Id, &user.FirstName, &user.SecondName, &user.Email, &user.Username, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("user repos: get all: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepos) GetOneBy(ctx context.Context) (models.User, error) {
	argsStr := []string{}

	ctxKeys := []interface{}{models.UserId, models.UserFirstName, models.UserSecondName, models.UserEmail, models.UserUsername, models.UserPassword}

	for _, ctxKey := range ctxKeys {
		ctxValue := ctx.Value(ctxKey)
		if ctxValue != nil {
			ctxKeyString := string(ctxKey.(models.UserCtx))
			ctxKeyString = strings.TrimPrefix(ctxKeyString, "user_")

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
		userTable,
		whereCondition,
	)

	var user models.User

	err := ur.db.QueryRowContext(ctx, query).Scan(&user.Id, &user.FirstName, &user.SecondName, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("user repos: get one by: %w", err)
	}

	return user, nil
}

func (ur *UserRepos) Update(ctx context.Context, userId uint, user models.UserUpdate) (models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (ur *UserRepos) Delete(ctx context.Context, userId uint) error {
	panic("not implemented") // TODO: Implement
}
