package param

type RegisterRequest struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required,len=11,numeric"`
	Password    string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	UserInfo UserInfo `json:"user"`
	Tokens   Tokens   `json:"tokens"`
}
