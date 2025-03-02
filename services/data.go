package services

import (
	"bechallenge/types"
	"encoding/json"
	"os"
	"path/filepath"
)

type data struct {
	users   types.Users
	actions types.Actions
}

// NewDataService creates a new instance of the Data Service
func NewDataService(path string) (data, error) {
	actionsFile := filepath.Join(path, "actions.json")
	usersFile := filepath.Join(path, "users.json")

	// read in the actions.json file and unmarshal it
	var actions types.Actions
	actionsBytes, err := os.ReadFile(actionsFile)
	if err != nil {
		return data{}, err
	}
	err = json.Unmarshal(actionsBytes, &actions)
	if err != nil {
		return data{}, err
	}

	// read in the users.json file and unmarshal it
	var users types.Users
	usersBytes, err := os.ReadFile(usersFile)
	if err != nil {
		return data{}, err
	}
	err = json.Unmarshal(usersBytes, &users)
	if err != nil {
		return data{}, err
	}

	//create a data struct with this unmarshalled data
	return data{
		users:   users,
		actions: actions,
	}, nil
}

// Users returns a slice of all users
func (d *data) Users() types.Users {
	return d.users
}

// Actions returns a slice of all actions
func (d *data) Actions() types.Actions {
	return d.actions
}
