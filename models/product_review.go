package models

type ProductReview struct {
	ID        int    `json:"id,omitempty"`
	UserId    int    `json:"userid,omitempty"`
	ProductId int    `json:"productid,omitempty"`
	Review    string `json:"review"`
	Rating    int    `json:"rating,omitempty"`
	Date      string `json:"Date,omitempty"`
}

type ProductReviewsResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    []ProductReview `json:"data,omitempty"`
}

type ProductReviewResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    ProductReview `json:"data,omitempty"`
}
