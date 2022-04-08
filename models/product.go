package models

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	StoreId  int    `json:"storeId"`
}

type ProductsResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Product `json:"data,omitempty"`
}

type ProductResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Product `json:"data,omitempty"`
}
