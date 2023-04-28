package entities

type Store struct {
	ID          int    `json:"id"`
	ExternalID  string `json:"external_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
