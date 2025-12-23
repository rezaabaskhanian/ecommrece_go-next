package entity

import "time"

type Cart struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Status    string    `json:"status"` // active, ordered, canceled
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
