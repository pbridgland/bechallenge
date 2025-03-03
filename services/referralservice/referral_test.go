package referralservice

import (
	"bechallenge/mocks"
	"errors"
	"reflect"
	"testing"
)

func TestReferralIndexes(t *testing.T) {
	mockService := mocks.DataRepo{}
	r := NewReferralService(&mockService)
	tests := []struct {
		name                    string
		mockSetup               func(t *testing.T)
		expectedReferralIndexes map[int]int
		expectedError           error
	}{
		{
			name: "Test Valid Simple Tree",
			mockSetup: func(t *testing.T) {
				err := mockService.SetSampleData("../../mocks/mockdata/referralTreeUsers.json", "../../mocks/mockdata/referralTreeActions.json")
				if err != nil {
					t.Fatalf("%v", err)
				}
			},
			expectedReferralIndexes: map[int]int{
				1:  9,
				2:  3,
				3:  3,
				4:  0,
				5:  1,
				6:  0,
				7:  0,
				8:  2,
				9:  1,
				10: 0,
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(t)

			gotReferralIndexes, gotErr := r.ReferralIndexes()

			if !errors.Is(gotErr, tt.expectedError) {
				t.Errorf("expected error to wrap %v but got %v", tt.expectedError, gotErr)
			}
			if tt.expectedError != nil {
				return
			}

			if !reflect.DeepEqual(tt.expectedReferralIndexes, gotReferralIndexes) {
				t.Errorf("expected referral indexes to be %v but got %v", tt.expectedReferralIndexes, gotReferralIndexes)
			}
		})
	}
}
