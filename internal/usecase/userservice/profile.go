package userservice

import (
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/param"
	"github.com/rezaabaskhanian/ecommrece_go-next.git/internal/pkg/richerror"
)

func (s Service) Profile(req param.ProfileRequest) (param.ProfileResponse, error) {
	const op = "userservice.profile"

	resp, err := s.repo.GetUserByID(int(req.UserID))

	if err != nil {
		return param.ProfileResponse{}, richerror.New(op).WithErr(err).WithMessage("dont get from GETUSERBYID")
	}

	return param.ProfileResponse{
		Name:        resp.Name,
		PhoneNumber: resp.PhoneNumber,
		AvatarURL: func() string {
			if resp.AvatarURL != nil {
				return *resp.AvatarURL
			}
			return ""
		}(),
	}, nil

}
