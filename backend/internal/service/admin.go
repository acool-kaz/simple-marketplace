package service

import (
	"errors"
	"main/internal/models"
	"main/internal/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

type Admin interface {
	SignIn(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type AdminService struct {
	repos repository.Admin
}

func newAdminService(repos repository.Admin) *AdminService {
	return &AdminService{
		repos: repos,
	}
}

func (s *AdminService) SignIn(username, password string) (string, error) {
	admin, err := s.repos.GetAdmin(username, password)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		admin.Id,
	})
	return token.SignedString([]byte(models.Key))
}

func (s *AdminService) ParseToken(accessToken string) (int, error) {
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
