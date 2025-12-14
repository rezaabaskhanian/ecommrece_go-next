package param

type PasswordRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type PasswordResponse struct {
	Password string `json:"password"`
}
