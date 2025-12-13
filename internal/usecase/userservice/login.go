package userservice

import (
	"fmt"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/errmsg"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	const op = "userservice.Login"

	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)

	if err != nil {
		return param.LoginResponse{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotFound)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return param.LoginResponse{}, richerror.New(op).WithErr(err).WithMessage("password is wrong")

	}

	accessToekn, err := s.auth.CreateAccessToken(user)

	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)

	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	return param.LoginResponse{
		UserInfo: param.UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
		Tokens: param.Tokens{
			AccessToken:  accessToekn,
			RefreshToken: refreshToken,
		},
	}, err

}
