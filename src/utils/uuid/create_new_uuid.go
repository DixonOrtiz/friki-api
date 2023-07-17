package uuid

import gouuid "github.com/google/uuid"

func New() string {
	uid := gouuid.New()
	return uid.String()
}
