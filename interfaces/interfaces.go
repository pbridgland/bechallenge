package interfaces

import "bechallenge/types"

// ProcessingService is used to get processed data for users and actions
type ProcessingService interface {
	User(id int) (types.User, error)
	UserActionCount(id int) (int, error)
	NextActions(actionType string) (map[string]float64, error)
}

// DataRepo is used to get raw data for users and actions
type DataRepo interface {
	Users() types.Users
	Actions() types.Actions
}

// ReferralService is used to get data relating to user referrals
type ReferralService interface {
	ReferralIndexes() (map[int]int, error)
}
