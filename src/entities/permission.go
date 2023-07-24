package entities

type Permission struct {
	ID        string   `json:"id"`
	UserID    string   `json:"user_id"`
	Addresses []string `json:"addresses"`
	Stores    []string `json:"stores"`
}
