package handlers

import (
	"bechallenge/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleNextActions(t *testing.T) {
	mockService := mocks.ProcessingService{}
	n := NewNextActionsHandler(&mockService)
	tests := []struct {
		name           string
		urlPath        string
		mockSetup      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Invalid action type format",
			urlPath:        "/actions//nextactions",
			mockSetup:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid action type\n",
		},
		{
			name:    "Internal server error",
			urlPath: "/actions/test/nextactions",
			mockSetup: func() {
				mockService.NextErrorForNextActions = errors.New("test error")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "error getting next actions for action type: test\n",
		},
		{
			name:    "Valid response",
			urlPath: "/actions/test/nextactions",
			mockSetup: func() {
				mockService.NextErrorForNextActions = nil
				mockService.NextResultForNextActions = map[string]float64{
					"act1": 0.6,
					"act2": 0.4,
				}
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"act1":0.6,"act2":0.4}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest(http.MethodGet, tt.urlPath, nil)
			w := httptest.NewRecorder()

			n.Handle(w, req)

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
