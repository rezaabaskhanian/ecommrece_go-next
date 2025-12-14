package userservice

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/errmsg"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) ResetPassword(req param.PasswordRequest) error {

	const op = "userservice.ResetPassword"

	//   passwordNew := func() string {
	//       if req.Password ==""{
	// 		return
	// 	  }
	//   }()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomthingWentWrong)
	}

	errRepo := s.repo.ResetPassword(req.PhoneNumber, string(hashedPassword))

	if errRepo != nil {
		return richerror.New(op).WithErr(err).WithMessage("failed to reset password")
	}

	return nil

}
