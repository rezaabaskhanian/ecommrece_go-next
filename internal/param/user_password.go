package param

type PasswordRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,len=11,numeric"`
	Password    string `json:"password" validate:"required"`
}

type PasswordResponse struct {
	Password string `json:"password"`
}
