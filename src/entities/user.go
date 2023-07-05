package entities

type User struct {
	ExternalID string `json:"external_id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Picture    string `json:"picture"`
}
