package http_types

type ApiResponse struct {
	OK         bool        `json:"ok"`
	StatusCode int         `json:"status_code"`
	Message    *string     `json:"message"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"errors"`
}
