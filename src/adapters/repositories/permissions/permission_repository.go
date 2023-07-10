package permrepo

import "cloud.google.com/go/firestore"

type PermissionRepository struct {
	DB *firestore.Client
}
