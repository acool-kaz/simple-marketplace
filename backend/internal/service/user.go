package service

import (
	"context"
	"fmt"

	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/acool-kaz/simple-marketplace/internal/repository"
)

type UserService struct {
	userRepos repository.User
}

func newUserService(userRepos repository.User) *UserService {
	return &UserService{
		userRepos: userRepos,
	}
}

func (us *UserService) Create(ctx context.Context, user models.UserSignUp) (uint, error) {
	id, err := us.userRepos.Create(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("user serivce: create: %w", err)
	}

	return id, err
}

func (us *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	users, err := us.userRepos.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("user service: get all: %w", err)
	}

	ctx = context.WithValue(ctx, models.UserId, 1)
	fmt.Println(us.userRepos.GetOneBy(ctx))

	return users, nil
}

func (us *UserService) GetOneBy(ctx context.Context) (models.User, error) {
	user, err := us.userRepos.GetOneBy(ctx)
	if err != nil {
		return models.User{}, fmt.Errorf("user service: get one by: %w", err)
	}

	return user, nil
}

func (us *UserService) Update(ctx context.Context, userId uint, user models.UserUpdate) (models.User, error) {
	err := us.userRepos.Update(ctx, userId, user)
	if err != nil {
		return models.User{}, fmt.Errorf("user service: update: %w", err)
	}

	newUser, err := us.userRepos.GetOneBy(context.WithValue(ctx, models.UserId, userId))
	if err != nil {
		return models.User{}, fmt.Errorf("user service: update: %w", err)
	}

	return newUser, nil
}

func (us *UserService) Delete(ctx context.Context, userId uint) error {
	err := us.userRepos.Delete(ctx, userId)
	if err != nil {
		return fmt.Errorf("user service: delete: %w", err)
	}

	return nil
}
