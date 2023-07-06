package types

import "time"

type FirestoreUser struct {
	ExternalID string    `firestore:"external_id"`
	Name       string    `firestore:"name"`
	LastName   string    `firestore:"last_name"`
	Email      string    `firestore:"email"`
	Picture    string    `firestore:"picture"`
	CreatedAt  time.Time `firestore:"created_at"`
	UpdatedAt  time.Time `firestore:"updated_at"`
}
