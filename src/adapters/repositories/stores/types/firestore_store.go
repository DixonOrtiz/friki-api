package types

import "time"

type FirestoreStore struct {
	ID          string    `firestore:"id"`
	UserID      string    `firestore:"user_id"`
	Name        string    `firestore:"name"`
	Description string    `firestore:"description"`
	CreatedAt   time.Time `firestore:"created_at"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}
