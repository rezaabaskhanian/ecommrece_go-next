package param

type PaginateResponse struct {
	Products    []Product `json:"products"`
	CurrentPage int       `json:"currentPage"`
	TotalPage   int       `json:"totalPage"`
	TotalItems  int       `json:"totalItems"`
}

type PaginateRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
