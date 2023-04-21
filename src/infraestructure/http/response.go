package httpinfra

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error any         `json:"error,omitempty"`
}
