package permrepo

import "cloud.google.com/go/firestore"

func MakePermissionRepository(DB *firestore.Client) IPermissionRepository {
	return &PermissionRepository{
		DB: DB,
	}
}
