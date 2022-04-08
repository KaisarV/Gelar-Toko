package models

type Chat struct {
	ID         int    `json:"id,omitempty"`
	SenderId   int    `json:"senderid,omitempty"`
	ReceiverId int    `json:"receiverid,omitempty"`
	Text       string `json:"text"`
	Date       string `json:"Date,omitempty"`
}

type ChatsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Chat `json:"data,omitempty"`
}

type ChatResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Chat   `json:"data,omitempty"`
}
