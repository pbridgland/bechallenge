package handlers

import (
	"bechallenge/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleReferralIndexes(t *testing.T) {
	mockService := mocks.ReferralService{}
	r := NewReferralIndexesHandler(&mockService)
	tests := []struct {
		name           string
		urlPath        string
		mockSetup      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "Internal server error",
			urlPath: "/referralindexes",
			mockSetup: func() {
				mockService.NextReferralIndexesErr = errors.New("test error")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "error getting referralindexes\n",
		},
		{
			name:    "Valid response",
			urlPath: "/referralindexes",
			mockSetup: func() {
				mockService.NextReferralIndexesErr = nil
				mockService.NextReferralIndexesResult = map[int]int{
					1: 2,
					3: 4,
				}
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"1":2,"3":4}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest(http.MethodGet, tt.urlPath, nil)
			w := httptest.NewRecorder()

			r.Handle(w, req)

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
