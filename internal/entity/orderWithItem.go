package entity

type OrderWithItem struct {
	Order Order       `json:"order"`
	Items []OrderItem `json:"items"`
}
