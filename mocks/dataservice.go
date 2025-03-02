package mocks

import "bechallenge/types"

type DataService struct {
	NextUsers   types.Users
	NextActions types.Actions
}

func (d *DataService) Users() types.Users {
	return d.NextUsers
}

func (d *DataService) Actions() types.Actions {
	return d.NextActions
}
