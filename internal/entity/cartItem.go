package entity

type CartItem struct {
	CartItemID int     `json:"cart_item_id"`
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	ImageURL   string  `json:"image_url,omitempty"`
}
