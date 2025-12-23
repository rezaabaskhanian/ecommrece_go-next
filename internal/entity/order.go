package entity

import "time"

type Order struct {
	ID        int       `json:"id"`
	USerID    int       `json:"user_id"`
	CartID    int       `json:"cart_id"`
	Status    string    `json:"status"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
