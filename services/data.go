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

func NewDataService(path string) (data, error) {
	actionsFile := filepath.Join(path, "actions.json")
	usersFile := filepath.Join(path, "users.json")
	var actions types.Actions
	actionsBytes, err := os.ReadFile(actionsFile)
	if err != nil {
		return data{}, err
	}
	err = json.Unmarshal(actionsBytes, &actions)
	if err != nil {
		return data{}, err
	}

	var users types.Users
	usersBytes, err := os.ReadFile(usersFile)
	if err != nil {
		return data{}, err
	}
	err = json.Unmarshal(usersBytes, &users)
	if err != nil {
		return data{}, err
	}

	return data{
		users:   users,
		actions: actions,
	}, nil
}

func (d *data) Users() types.Users {
	return d.users
}

func (d *data) Actions() types.Actions {
	return d.actions
}
