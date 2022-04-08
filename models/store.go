package models

type Store struct {
	ID      int    `json:"id"`
	UserId  int    `json:"userId,omitempty"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type StoresResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Store `json:"data,omitempty"`
}

type StoreResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Store  `json:"data,omitempty"`
}
