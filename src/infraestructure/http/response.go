package httpinfra

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   any         `json:"error,omitempty"`
}
