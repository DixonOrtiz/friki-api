package entities

import "time"

type Address struct {
	ID               string    `json:"id"` // este hay que generarle
	UserID           string    `json:"user_id"`
	Name             string    `json:"name"`
	City             string    `json:"city"`
	Commune          string    `json:"commune"`
	Street           string    `json:"street"`
	Number           int       `json:"number"`
	Apartment        string    `json:"apartment"`
	Sector           string    `json:"sector"`
	Type             string    `json:"type"`
	ExtraInformation string    `json:"extra_information"`
	CreatedAt        time.Time `json:"created_at"` // se crea
	UpdatedAt        time.Time `json:"updated_at"` // se crea
}
