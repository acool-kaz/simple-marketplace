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
	panic("not implemented") // TODO: Implement
}

func (us *UserService) Update(ctx context.Context, userId uint, user models.UserUpdate) (models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (us *UserService) Delete(ctx context.Context, userId uint) error {
	panic("not implemented") // TODO: Implement
}
