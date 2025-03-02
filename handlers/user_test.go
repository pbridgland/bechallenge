package handlers

import (
	"bechallenge/mocks"
	"bechallenge/types"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandle(t *testing.T) {
	mockService := mocks.ProcessingService{}
	u := NewUserHandler(&mockService)
	tests := []struct {
		name           string
		urlPath        string
		mockSetup      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Invalid user ID format",
			urlPath:        "/users/abc",
			mockSetup:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid user ID\n",
		},
		{
			name:    "User not found",
			urlPath: "/users/42",
			mockSetup: func() {
				mockService.NextUserErr = types.ErrUserNotPresent
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 page not found\n",
		},
		{
			name:    "Internal server error",
			urlPath: "/users/99",
			mockSetup: func() {
				mockService.NextUserErr = errors.New("test error")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "error getting user with ID: 99\n",
		},
		{
			name:    "Valid user ID",
			urlPath: "/users/1",
			mockSetup: func() {
				mockService.NextUserErr = nil
				mockService.NextUserResult = types.User{
					ID:        1,
					Name:      "Test Name",
					CreatedAt: time.Date(2025, 03, 02, 11, 37, 0, 0, time.UTC),
				}
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"name":"Test Name","createdAt":"2025-03-02T11:37:00Z"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest(http.MethodGet, tt.urlPath, nil)
			w := httptest.NewRecorder()

			u.Handle(w, req)

			res := w.Result()
			defer res.Body.Close()

			if tt.expectedStatus != res.StatusCode {
				t.Errorf("expected status code to be %d but got %d", tt.expectedStatus, res.StatusCode)
			}

			bodyBytes := w.Body.Bytes()
			bodyString := string(bodyBytes)
			if tt.expectedBody != bodyString {
				t.Errorf("expected body to be %s but got %s", tt.expectedBody, bodyString)
			}
		})
	}
}
