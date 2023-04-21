package contports

import "frikiapi/src/entities"

type UserUseCases interface {
	AuthorizeRole(userID int, roles []int) (entities.User, error)
}
