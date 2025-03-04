package types

import (
	"errors"
	"fmt"
	"time"
)

var ErrUserNotPresent = errors.New("user not present")

type Users []User
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// WithID returns the user with the given ID
func (u Users) WithID(id int) (User, error) {
	for _, user := range u {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("could not find user with id in slice. %w", ErrUserNotPresent)
}

type Actions []Action
type Action struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	UserID     int       `json:"userID"`     // The ID of the User who performed this action
	TargetUser int       `json:"targetUser"` // Supplied when "REFER_USER" action type
	CreatedAt  time.Time `json:"createdAt"`
}

// ByUserWithID returns all actions that were performed by the user with the given ID
func (a Actions) ByUserWithID(userID int) Actions {
	actionsByUser := make(Actions, 0)
	for _, action := range a {
		if action.UserID == userID {
			actionsByUser = append(actionsByUser, action)
		}
	}
	return actionsByUser
}
