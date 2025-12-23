package entity

type CartWithItem struct {
	Cart       Cart
	Items      []CartItem
	TotalPrice float64
	ItemCount  int
}
