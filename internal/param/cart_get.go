package param

type CreateGetRequest struct {
	UserID int `json:"user_id"`
}

type CartGetResposne struct {
	Cart CartInfo `json:"cart_info"`
}
