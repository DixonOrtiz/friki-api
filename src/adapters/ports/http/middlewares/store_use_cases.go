package midports

type StoreUseCases interface {
	Authorize(externalID string, storeID int) error
}
