package authservice

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

type Config struct {
	SignKey               string
	AccessExpirationTime  time.Duration
	RefreshExpiratoonTime time.Duration

	AccessSubject  string
	RefreshSubject string
}

type Service struct {
	config Config
}

func New(cfg Config) Service {

	return Service{
		config: cfg,
	}
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.creatoken(user.ID, user.Role, s.config.AccessSubject, s.config.AccessExpirationTime)

}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.creatoken(user.ID, user.Role, s.config.RefreshSubject, s.config.RefreshExpiratoonTime)

}

func (s Service) ParseToken(authHeader string) (*Claims, error) {
	const op = "authservice.parseToken"

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, richerror.New(op)
	}

	tokenStr := parts[1]

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, richerror.New(op).WithMessage("invalid token claims")

}

func (s Service) creatoken(userID uint, role, subject string, expireDuration time.Duration) (string, error) {
	const op = "auhtservice.createtoken"
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: subject,
			//TODO: set the expire time
			//TODO: see https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := accessToken.SignedString([]byte(s.config.SignKey))

	if err != nil {
		return "", richerror.New(op).WithErr(err).WithMessage("dont create token")
	}

	return tokenString, nil

}
