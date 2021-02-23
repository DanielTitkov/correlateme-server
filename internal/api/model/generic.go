package model

type (
	// OKResponse is default ok response schema
	OKResponse struct {
		Status  string `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
	}
	// ErrorResponse is defailt error response schema
	ErrorResponse struct {
		Error   string `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
	}
)
