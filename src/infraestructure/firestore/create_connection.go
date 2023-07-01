package firestore

import (
	"context"
	"os"

	gofirestore "cloud.google.com/go/firestore"
)

func CreateConnection() (*gofirestore.Client, error) {
	client, err := gofirestore.NewClient(context.Background(), os.Getenv("PROJECT_ID"))
	if err != nil {
		return nil, err
	}

	return client, nil
}