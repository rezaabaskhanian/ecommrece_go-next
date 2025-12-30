package param

type ProductWithCategory struct {
	ID           int
	Name         string
	Description  string
	Price        float64
	Stock        int
	ImageURL     string
	CategoryID   int
	CategoryName string
	CategorySlug string
}
