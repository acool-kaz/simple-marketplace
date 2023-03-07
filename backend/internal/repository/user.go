package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/acool-kaz/simple-marketplace/internal/models"
	sortfilter "github.com/acool-kaz/simple-marketplace/pkg/sort_filter"
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
			(first_name, second_name, email, phone_number, username, password) 
		VALUES 
			('%s', '%s', '%s', '%s', '%s', '%s') 
		RETURNING id;`,
		userTable,
		user.FirstName, user.SecondName, user.Email, user.PhoneNumber, user.Username, user.Password,
	)

	var id uint
	err := ur.db.QueryRowContext(ctx, query).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("user repos: create: %w", err)
	}

	return id, err
}

func (ur *UserRepos) GetAll(ctx context.Context) ([]models.User, error) {
	sortByQuery := ""

	sortBy := ctx.Value(models.UserSortBy)
	if sortBy != nil {
		sort, err := sortfilter.ValidateAndReturnSortQuery(sortBy.(string), models.UserSortParams)
		if err != nil {
			return nil, fmt.Errorf("user repos: get all: %w", err)
		}

		sortByQuery = sort
	}

	query := fmt.Sprintf(`
		SELECT
			*
		FROM %s
		%s;`,
		userTable,
		sortByQuery,
	)

	rows, err := ur.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("user repos: get all: %w", err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.Id, &user.Role, &user.FirstName, &user.SecondName, &user.Email, &user.PhoneNumber, &user.Username, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("user repos: get all: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepos) GetOneBy(ctx context.Context) (models.User, error) {
	argsStr := []string{}

	ctxKeys := []interface{}{models.UserId, models.UserRole, models.UserFirstName, models.UserSecondName, models.UserEmail, models.UserPhoneNumber, models.UserUsername, models.UserPassword}

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

	err := ur.db.QueryRowContext(ctx, query).Scan(&user.Id, &user.Role, &user.FirstName, &user.SecondName, &user.Email, &user.PhoneNumber, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = models.ErrUserNotFound
		}
		return models.User{}, fmt.Errorf("user repos: get one by: %w", err)
	}

	return user, nil
}

func (ur *UserRepos) Update(ctx context.Context, userId uint, user models.UserUpdate) error {
	argsStr := []string{}

	if user.FirstName != "" {
		argsStr = append(argsStr, fmt.Sprintf("first_name = '%s'", user.FirstName))
	}

	if user.SecondName != "" {
		argsStr = append(argsStr, fmt.Sprintf("second_name = '%s'", user.SecondName))
	}

	if user.Email != "" {
		argsStr = append(argsStr, fmt.Sprintf("email = '%s'", user.Email))
	}

	if user.PhoneNumber != "" {
		argsStr = append(argsStr, fmt.Sprintf("phone_number = '%s'", user.PhoneNumber))
	}

	if user.Username != "" {
		argsStr = append(argsStr, fmt.Sprintf("username = '%s'", user.Username))
	}

	if user.Password != "" {
		argsStr = append(argsStr, fmt.Sprintf("password = '%s'", user.Password))
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
		userId,
	)

	_, err := ur.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("user repos: update: %w", err)
	}

	return nil
}

func (ur *UserRepos) Delete(ctx context.Context, userId uint) error {
	query := fmt.Sprintf(`
		DELETE FROM
			%s
		WHERE id = '%d'`,
		userTable,
		userId,
	)

	_, err := ur.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("user repos: delete: %w", err)
	}

	return nil
}
