package models

type Cart struct {
	ID        int `json:"id"`
	UserId    int `json:"userid"`
	ProductId int `json:"ProductId"`
	Quantity  int `json:"qty"`
}

type CartsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Cart `json:"data,omitempty"`
}

type CartResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Cart   `json:"data,omitempty"`
}
