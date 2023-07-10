package types

import "time"

type FirestoreAddress struct {
	ID               string    `firestore:"id"`
	UserID           string    `firestore:"user_id"`
	Name             string    `firestore:"name"`
	City             string    `firestore:"city"`
	Commune          string    `firestore:"commune"`
	Street           string    `firestore:"street"`
	Number           int       `firestore:"number"`
	Apartment        string    `firestore:"apartment"`
	Sector           string    `firestore:"sector"`
	Type             string    `firestore:"type"`
	ExtraInformation string    `firestore:"extra_information"`
	CreatedAt        time.Time `firestore:"created_at"`
	UpdatedAt        time.Time `firestore:"updated_at"`
}
