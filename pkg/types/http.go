package types

type ApiResponse struct {
	OK         bool        `json:"ok"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"errors"`
}
