package param

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,len=11,numeric"`
	Password    string `json:"password" validate:"required"`
}

type LoginResponse struct {
	UserInfo UserInfo `json:"user"`
	Tokens   Tokens   `json:"tokens"`
}
