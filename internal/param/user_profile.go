package param

type ProfileRequest struct {
	UserID uint `json:"id"`
}

type ProfileResponse struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	AvatarURL   string `json:"avatar_url"`
}
