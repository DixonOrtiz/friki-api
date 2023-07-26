package types

type FirestorePermission struct {
	ID        string   `firestore:"id"`
	UserID    string   `firestore:"user_id"`
	Addresses []string `firestore:"addresses"`
	Stores    []string `firestore:"stores"`
}
