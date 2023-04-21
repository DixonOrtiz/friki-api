package rrutils

type RequestRoleInput struct {
	Role   int
	Reason string
}

func (i *RequestRoleInput) GetValues() (int, string) {
	return i.Role, i.Reason
}
