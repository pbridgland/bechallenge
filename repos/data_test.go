package repos

import (
	"bechallenge/types"
	"reflect"
	"testing"
	"time"
)

func TestData(t *testing.T) {
	d, err := NewDataRepo("../data")
	if err != nil {
		t.Errorf("expected err to be nil, got %v", err)
	}
	testUsers := types.Users{
		types.User{
			ID:        1,
			Name:      "TestName",
			CreatedAt: time.Date(2025, 03, 02, 11, 37, 0, 0, time.UTC),
		},
	}
	d.users = testUsers
	gotUsers := d.Users()
	if !reflect.DeepEqual(gotUsers, testUsers) {
		t.Errorf("expected calling Users to return %v got %v", testUsers, gotUsers)
	}

	testActions := types.Actions{
		types.Action{
			ID:         1,
			Type:       "Test Type",
			UserID:     2,
			TargetUser: 3,
			CreatedAt:  time.Date(2025, 03, 02, 11, 37, 0, 0, time.UTC),
		},
	}
	d.actions = testActions
	gotActions := d.Actions()
	if !reflect.DeepEqual(gotActions, testActions) {
		t.Errorf("expected calling Actions to return %v got %v", testActions, gotActions)
	}
}
