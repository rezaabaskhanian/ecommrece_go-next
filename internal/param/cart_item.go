package param

type CartItemRequest struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CartItemResponse struct {
	CartItemID int     `json:"cart_item_id"`
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	ImageURL   string  `json:"image_url,omitempty"`
}
