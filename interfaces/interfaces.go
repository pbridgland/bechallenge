package interfaces

import "bechallenge/types"

// ProcessingService specifies the methods a processing service must implement
type ProcessingService interface {
	User(id int) (types.User, error)
	UserActionCount(id int) (int, error)
	NextActions(actionType string) (map[string]float64, error)
}

// DataService specifies the methods a data service must implement
type DataService interface {
	Users() types.Users
	Actions() types.Actions
}
