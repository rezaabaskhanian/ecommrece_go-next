package param

type CartUpdateQuantityRequest struct {
	CartItemID int    `json:"cart_item_id"`
	Action     string `json:"action"`
}
