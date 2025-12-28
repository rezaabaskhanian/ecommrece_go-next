package param

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int    `json:"stock"`
	CategoryID  int    `json:"category_id"`
	ImageURL    string `json:"image_url"`
}
