package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	jwt.StandardClaims
	Id   uint
	Role string
}

const (
	AccessTokenTime  = 30 * 24 * time.Hour
	RefreshTokenTime = 30 * 24 * time.Hour
)
