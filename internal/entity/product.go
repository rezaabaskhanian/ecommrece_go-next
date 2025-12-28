package entity

import "time"

type Product struct {
	ID          int       `json:"id"`
	ShopID      int       `json:"shop_id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Price       int64     `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  int       `json:"category,omitempty"`
	ImageURL    string    `json:"image_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
