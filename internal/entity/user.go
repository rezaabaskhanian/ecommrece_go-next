package entity

import "time"

type User struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AvatarURL string `json:"avatar_url"`

	Role string `json:"role"` // user or admin

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
