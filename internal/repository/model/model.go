package model

type ProductWithCategory struct {
	ID           int
	Name         string
	Description  string
	Price        int64
	Stock        int
	ImageURL     string
	CategoryID   int
	CategoryName string
	CategorySlug string
}
