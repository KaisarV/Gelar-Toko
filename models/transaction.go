package models

type Transaction struct {
	ID        int    `json:"id"`
	UserId    int    `json:"userId"`
	ProductId int    `json:"productId"`
	Date      string `json:"date"`
	Quantity  int    `json:"qty"`
	Status    string `json:"status"`
}

type TransactionsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Transaction `json:"data,omitempty"`
}

type TransactionResponse struct {
	Status  int         `json:"status" gorm:"primary_key"`
	Message string      `json:"message"`
	Data    Transaction `json:"data,omitempty"`
}
