package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/acool-kaz/simple-marketplace/internal/models"
	"github.com/acool-kaz/simple-marketplace/internal/repository"
	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	userRepos repository.User
}

func newAuthService(userRepos repository.User) *AuthService {
	return &AuthService{
		userRepos: userRepos,
	}
}

func (as *AuthService) SignUp(ctx context.Context, user models.UserSignUp) (uint, error) {
	_, err := as.userRepos.GetOneBy(context.WithValue(ctx, models.UserEmail, user.Email))
	if err != nil && !errors.Is(err, models.ErrUserNotFound) {
		return 0, fmt.Errorf("auth service: sign up: %w", err)
	} else if err == nil {
		return 0, fmt.Errorf("auth service: sign up: %w", models.ErrUserEmailExist)
	}

	_, err = as.userRepos.GetOneBy(context.WithValue(ctx, models.UserUsername, user.Username))
	if err != nil && !errors.Is(err, models.ErrUserNotFound) {
		return 0, fmt.Errorf("auth service: sign up: %w", err)
	} else if err == nil {
		return 0, fmt.Errorf("auth service: sign up: %w", models.ErrUserUsernameExist)
	}

	id, err := as.userRepos.Create(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("auth service: sign up: %w", err)
	}

	return id, nil
}

func (as *AuthService) SignIn(ctx context.Context, user models.UserSignIn) (string, string, error) {
	var (
		curUser models.User
		err     error
	)

	if user.Email != "" {
		curUser, err = as.userRepos.GetOneBy(context.WithValue(ctx, models.UserEmail, user.Email))
	} else if user.Username != "" {
		curUser, err = as.userRepos.GetOneBy(context.WithValue(ctx, models.UserUsername, user.Username))
	}

	if err != nil {
		return "", "", fmt.Errorf("auth service: sign in: %w", err)
	}

	if curUser.Password != user.Password {
		return "", "", fmt.Errorf("auth service: sign in: %w", models.ErrUserNotFound)
	}

	access, err := newAccessToken(curUser.Id, curUser.Role)
	if err != nil {
		return "", "", fmt.Errorf("auth service: sign in: %w", err)
	}

	refresh, err := newRefreshToken(curUser.Id, curUser.Role)
	if err != nil {
		return "", "", fmt.Errorf("auth service: sign in: %w", err)
	}

	return access, refresh, nil
}

func (as *AuthService) ParseToken(ctx context.Context, accessToken string) (*models.Token, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.Token{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth service: parse token: invalid signing method")
		}
		return []byte(os.Getenv("JWT_SALT")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("auth service: parse token: %w", err)
	}

	claims, ok := token.Claims.(*models.Token)
	if !ok {
		return nil, errors.New("auth service: parse token: token claims are not of type *tokenClaims")
	}

	return claims, nil
}

func newAccessToken(userId uint, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Token{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(models.AccessTokenTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:   userId,
		Role: role,
	})

	return token.SignedString([]byte(os.Getenv("JWT_SALT")))
}

func newRefreshToken(userId uint, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Token{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(models.RefreshTokenTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:   userId,
		Role: role,
	})

	return token.SignedString([]byte(os.Getenv("JWT_SALT")))
}
