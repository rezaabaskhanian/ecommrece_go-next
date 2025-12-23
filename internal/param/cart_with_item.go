package param

type CartWithItem struct {
	Cart       CartInfo           `json:"cart_info"`
	Items      []CartItemResponse `json:"cart_item"`
	TotalPrice float64            `json:"total_price"`
	ItemCount  int                `json:"item_count"`
}
