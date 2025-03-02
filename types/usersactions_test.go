package types

import (
	"errors"
	"reflect"
	"testing"
)

func TestWithID(t *testing.T) {
	tests := []struct {
		name          string
		users         Users
		id            int
		expectedUser  User
		expectedError error
	}{
		{
			name: "User found",
			id:   1,
			users: Users{
				User{ID: 1},
			},
			expectedUser:  User{ID: 1},
			expectedError: nil,
		},
		{
			name: "User not found",
			id:   1,
			users: Users{
				User{ID: 1},
			},
			expectedUser:  User{ID: 1},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, gotErr := tt.users.WithID(tt.id)

			if !errors.Is(gotErr, tt.expectedError) {
				t.Errorf("expected error to wrap %v but got %v", tt.expectedError, gotErr)
			}
			if tt.expectedError != nil {
				return
			}

			if tt.expectedUser != gotUser {
				t.Errorf("expected user to be %v but got %v", tt.expectedUser, gotUser)
			}
		})
	}
}

func TestByUserWithID(t *testing.T) {
	tests := []struct {
		name            string
		actions         Actions
		id              int
		expectedActions Actions
	}{
		{
			name: "Actions counted",
			id:   1,
			actions: Actions{
				Action{
					ID:     1,
					UserID: 1,
				},
				Action{
					ID:     2,
					UserID: 2,
				},
				Action{
					ID:     3,
					UserID: 1,
				},
				Action{
					ID:     4,
					UserID: 3,
				},
			},
			expectedActions: Actions{
				Action{
					ID:     1,
					UserID: 1,
				},
				Action{
					ID:     3,
					UserID: 1,
				},
			},
		},
		{
			name: "0 Actions matching",
			id:   7,
			actions: Actions{
				Action{
					ID:     1,
					UserID: 1,
				},
				Action{
					ID:     2,
					UserID: 2,
				},
				Action{
					ID:     3,
					UserID: 1,
				},
				Action{
					ID:     4,
					UserID: 3,
				},
			},
			expectedActions: Actions{},
		},
		{
			name:            "no actions",
			id:              7,
			actions:         Actions{},
			expectedActions: Actions{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotActions := tt.actions.ByUserWithID(tt.id)

			if !reflect.DeepEqual(tt.expectedActions, gotActions) {
				t.Errorf("expected count to be %v but got %v", tt.expectedActions, gotActions)
			}
		})
	}
}
