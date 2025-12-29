package authservice

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID uint   `json:"UserID"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
