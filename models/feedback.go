package models

type Feedback struct {
	ID       int    `json:"id,omitempty"`
	UserId   int    `json:"userid,omitempty"`
	Feedback string `json:"feedback"`
	Date     string `json:"Date,omitempty"`
}

type FeedbacksResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Feedback `json:"data,omitempty"`
}

type FeedbackResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    Feedback `json:"data,omitempty"`
}
