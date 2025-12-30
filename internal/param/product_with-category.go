package param

type ProductWithCategoryReq struct {
	Slug  string `json:"slug"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}
