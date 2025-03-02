package handlers

import (
	"bechallenge/mocks"
	"bechallenge/types"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleUserActionCount(t *testing.T) {
	mockService := mocks.ProcessingService{}
	uac := NewUserActionCountHandler(&mockService)
	tests := []struct {
		name           string
		urlPath        string
		mockSetup      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Invalid user ID format",
			urlPath:        "/users/abc/actions/count",
			mockSetup:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid user ID\n",
		},
		{
			name:    "User not found",
			urlPath: "/users/42/actions/count",
			mockSetup: func() {
				mockService.NextUserActionCountErr = types.ErrUserNotPresent
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 page not found\n",
		},
		{
			name:    "Internal server error",
			urlPath: "/users/99/actions/count",
			mockSetup: func() {
				mockService.NextUserActionCountErr = errors.New("test error")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "error getting action count for user with ID: 99\n",
		},
		{
			name:    "Valid user ID",
			urlPath: "/users/1/actions/count",
			mockSetup: func() {
				mockService.NextUserActionCountErr = nil
				mockService.NextUserActionCountResult = 5
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"count":5}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest(http.MethodGet, tt.urlPath, nil)
			w := httptest.NewRecorder()

			uac.Handle(w, req)

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
