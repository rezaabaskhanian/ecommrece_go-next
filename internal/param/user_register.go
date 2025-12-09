package param

type RegisterRequest struct {
	Name        string
	PhoneNumber string
	Password    string
}

type RegisterResponse struct {
	UserInfo UserInfo `json:"user"`
}
