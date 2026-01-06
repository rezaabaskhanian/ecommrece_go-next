package userservice

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {
	const op = "usecase.register"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(op, "password hashing failed:", err)
		return param.RegisterResponse{}, fmt.Errorf("password hashing failed")
	}

	user := entity.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    string(hashedPassword),
		CreatedAt:   time.Now(),
	}

	createuser, err := s.repo.Register(user)
	if err != nil {
		log.Println(op, "db insert failed:", err)
		if strings.Contains(err.Error(), "duplicate") {
			return param.RegisterResponse{}, fmt.Errorf("phone number already registered")
		}
		return param.RegisterResponse{}, fmt.Errorf("database insert failed")
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		log.Println(op, "access token creation failed:", err)
		return param.RegisterResponse{}, fmt.Errorf("unexpected error creating access token")
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		log.Println(op, "refresh token creation failed:", err)
		return param.RegisterResponse{}, fmt.Errorf("unexpected error creating refresh token")
	}

	return param.RegisterResponse{
		UserInfo: param.UserInfo{
			ID:          createuser.ID,
			Name:        createuser.Name,
			PhoneNumber: createuser.PhoneNumber,
		},
		Tokens: param.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
