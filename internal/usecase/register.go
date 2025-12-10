package usecasse

import (
	"time"

	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/entity"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/errmsg"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {

	const op = "usecase.register"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return param.RegisterResponse{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomthingWentWrong)
	}

	user := entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    string(hashedPassword),
		CreatedAt:   time.Now(),
	}

	createuser, err := s.repo.Register(user)
	if err != nil {
		return param.RegisterResponse{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgSomthingWentWrong)
	}

	return param.RegisterResponse{
		// USerInfo write for better read
		UserInfo: param.UserInfo{
			ID:          createuser.ID,
			Name:        createuser.Name,
			PhoneNumber: createuser.PhoneNumber,
		},
	}, nil

}
