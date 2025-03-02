package processingservice

import (
	"bechallenge/mocks"
	"bechallenge/types"
	"errors"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	mockService := mocks.DataService{}
	p := NewProcessingService(&mockService)
	tests := []struct {
		name          string
		id            int
		mockSetup     func()
		expectedUser  types.User
		expectedError error
	}{
		{
			name: "Invalid user ID",
			id:   11,
			mockSetup: func() {
				mockService.NextUsers = types.Users{
					types.User{
						ID:        1,
						Name:      "Test Name",
						CreatedAt: time.Date(2025, 03, 02, 11, 37, 0, 0, time.UTC),
					},
				}
			},
			expectedUser:  types.User{},
			expectedError: types.ErrUserNotPresent,
		},
		{
			name: "Valid user ID",
			id:   1,
			mockSetup: func() {
				mockService.NextUsers = types.Users{
					types.User{
						ID:        1,
						Name:      "Test Name",
						CreatedAt: time.Date(2025, 03, 02, 11, 37, 0, 0, time.UTC),
					},
				}
			},
			expectedUser: types.User{
				ID:        1,
				Name:      "Test Name",
				CreatedAt: time.Date(2025, 03, 02, 11, 37, 0, 0, time.UTC),
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			gotUser, gotErr := p.User(tt.id)

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

func TestUserActionCount(t *testing.T) {
	mockService := mocks.DataService{}
	p := NewProcessingService(&mockService)
	tests := []struct {
		name          string
		id            int
		mockSetup     func(t *testing.T)
		expectedCount int
		expectedError error
	}{
		{
			name: "User with actions",
			id:   1,
			mockSetup: func(t *testing.T) {
				err := mockService.SetSampleData("../../mocks/mockdata/referralTreeUsers.json", "../../mocks/mockdata/referralTreeActions.json")
				if err != nil {
					t.Fatalf("%v", err)
				}
			},
			expectedCount: 3,
			expectedError: nil,
		},
		{
			name: "User with no actions",
			id:   10,
			mockSetup: func(t *testing.T) {
				err := mockService.SetSampleData("../../mocks/mockdata/referralTreeUsers.json", "../../mocks/mockdata/referralTreeActions.json")
				if err != nil {
					t.Fatalf("%v", err)
				}
			},
			expectedCount: 0,
			expectedError: nil,
		},
		{
			name: "Invalid UserID",
			id:   100,
			mockSetup: func(t *testing.T) {
				err := mockService.SetSampleData("../../mocks/mockdata/referralTreeUsers.json", "../../mocks/mockdata/referralTreeActions.json")
				if err != nil {
					t.Fatalf("%v", err)
				}
			},
			expectedCount: 0,
			expectedError: types.ErrUserNotPresent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(t)

			gotCount, gotErr := p.UserActionCount(tt.id)

			if !errors.Is(gotErr, tt.expectedError) {
				t.Errorf("expected error to wrap %v but got %v", tt.expectedError, gotErr)
			}
			if tt.expectedError != nil {
				return
			}

			if tt.expectedCount != gotCount {
				t.Errorf("expected count to be %d but got %d", tt.expectedCount, gotCount)
			}
		})
	}
}
