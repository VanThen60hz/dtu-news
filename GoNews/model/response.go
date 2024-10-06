package model

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Duration   string      `json:"duration,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}
