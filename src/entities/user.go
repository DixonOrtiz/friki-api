package entities

import "time"

type User struct {
	ExternalID string    `json:"external_id"`
	Name       string    `json:"name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Picture    string    `json:"picture"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
