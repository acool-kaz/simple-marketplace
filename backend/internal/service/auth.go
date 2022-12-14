package service

import (
	"errors"
	"main/internal/models"
	"main/internal/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

type Auth interface {
	SignUp(user models.User) error
	SignIn(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type AuthService struct {
	repos repository.Auth
}

func newAuthService(r repository.Auth) *AuthService {
	return &AuthService{
		repos: r,
	}
}

func (s *AuthService) SignUp(user models.User) error {
	return s.repos.CreateUser(user)
}

func (s *AuthService) SignIn(username, password string) (string, error) {
	user, err := s.repos.GetUser(username, password)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(models.Key))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(models.Key), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}
