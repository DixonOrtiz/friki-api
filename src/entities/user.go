package entities

type User struct {
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	ExternalID string `json:"external_id"`
	Picture    string `json:"picture"`
}
