package entity

import "time"

type User struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`

	AvatarURL *string `json:"avatar_url"` // nullable while use star use null into string

	Role string `json:"role"` // user or admin

	CreatedAt time.Time `json:"created_at"`
}
