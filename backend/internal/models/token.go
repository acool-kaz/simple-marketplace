package models

import "github.com/golang-jwt/jwt"

const Key = "kjHASkdHAIRsadkjahfhuH12390"

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
