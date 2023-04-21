package erutils

type EvaluateRoleInput struct {
	RoleRequestID int
	Approve       bool
}

func (i *EvaluateRoleInput) GetValues() (int, bool) {
	return i.RoleRequestID, i.Approve
}
