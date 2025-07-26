package handler

// JSON response handler format
type Response struct {
	Success bool   `json:"success"`
	Data    []any  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
