package addressusecases

func (u AddressUseCases) GetByID(tokenUserID, userID string, resources map[string]string) error {
	return u.PermissionUseCases.Authorize(tokenUserID, userID, resources)
}
