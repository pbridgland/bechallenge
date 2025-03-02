package interfaces

import "bechallenge/types"

type ProcessingService interface {
	User(id int) (types.User, error)
}

type DataService interface {
	Users() types.Users
	Actions() types.Actions
}
