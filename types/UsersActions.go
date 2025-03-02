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

type Actions []Action
type Action struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	UserID     int       `json:"userID"`     // The ID of the User who performed this action
	TargetUser int       `json:"targetUser"` // Supplied when "REFER_USER" action type
	CreatedAt  time.Time `json:"createdAt"`
}

func (u Users) WithID(id int) (User, error) {
	for _, user := range u {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("could not find user with id in slice. %w", ErrUserNotPresent)
}
