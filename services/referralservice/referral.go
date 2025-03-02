package referralservice

import (
	"bechallenge/interfaces"
	"bechallenge/types"
	"errors"
	"fmt"
)

const referUserType = "REFER_USER"

var err_methodCallOnNilReferralUser = errors.New("method called on nil referralUser")

type referral struct {
	data interfaces.DataService
}

// NewReferralService creates a new instance of the Processing Service
func NewReferralService(data interfaces.DataService) referral {
	return referral{
		data: data,
	}
}

// ReferralIndexes gets a map of userIDs to their referral Index
// complexity is O(n + m), where n is the number of users and m is the number of actions
// it loops over each action once, to find which users have referred which others
// it loops over users up to 3 times:
// - to create a "referralUser"
// - to get it's "referralIndex" for the results map
// - to get it's "referralIndex" if refered by another user, this will happen at most once as it is only calculated once and users can only be refered once
func (rs *referral) ReferralIndexes() (map[int]int, error) {
	users := rs.data.Users()
	actions := rs.data.Actions()

	referralUsersByID := referralUsersByID(users, actions)

	results := make(map[int]int, len(referralUsersByID))
	var err error
	// loop over each referralUser
	for id, user := range referralUsersByID {
		results[id], err = user.referralIndex() //calculate their referral index and store in the results map
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

// referralUsersByID creates a map of userIDs to referralUsers
func referralUsersByID(users types.Users, actions types.Actions) map[int]*referralUser {
	referralUsersByID := make(map[int]*referralUser, len(users))
	for _, user := range users {
		referralUsersByID[user.ID] = &referralUser{
			User: user,
		}
	}

	for _, action := range actions {
		if action.Type != referUserType || action.UserID == action.TargetUser {
			continue
		}
		parentUser := referralUsersByID[action.UserID]
		childUser := referralUsersByID[action.TargetUser]
		parentUser.addChild(childUser)
	}
	return referralUsersByID
}

// referralUser embeds the User type and gives extra data to aid calculating referral index
type referralUser struct {
	types.User
	children                []*referralUser // a pointer to each of the users this user has referred
	calculatedReferralIndex bool            // a flag to see if this user's referral index has been calcualted already
	refIndex                int
}

// addChild adds a pointer to the given referral user to the recievers slice of children
func (ru *referralUser) addChild(child *referralUser) {
	ru.children = append(ru.children, child)
}

// referralIndex calculates the referralIndex of the referralUser and stores the result
// if it is called again the stored result is returned
func (ru *referralUser) referralIndex() (int, error) {
	if ru == nil {
		return 0, fmt.Errorf("%w. on method referralIndex", err_methodCallOnNilReferralUser)
	}
	if ru.calculatedReferralIndex {
		return ru.refIndex, nil
	}
	index := 0
	for _, child := range ru.children {
		childReferralIndex, err := child.referralIndex()
		if err != nil {
			return 0, err
		}
		index += childReferralIndex + 1
	}
	ru.calculatedReferralIndex = true
	ru.refIndex = index
	return index, nil
}
