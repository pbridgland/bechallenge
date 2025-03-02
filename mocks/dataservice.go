package mocks

import (
	"bechallenge/types"
	"encoding/json"
	"os"
)

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

func (d *DataService) SetSampleData(userPath, actionPath string) error {
	var actions types.Actions
	actionsBytes, err := os.ReadFile(actionPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(actionsBytes, &actions)
	if err != nil {
		return err
	}

	var users types.Users
	usersBytes, err := os.ReadFile(userPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(usersBytes, &users)
	if err != nil {
		return err
	}

	d.NextUsers = users
	d.NextActions = actions

	return nil
}
