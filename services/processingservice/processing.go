package processingservice

import (
	"bechallenge/interfaces"
	"bechallenge/types"
	"math"
	"slices"
)

type processing struct {
	data interfaces.DataRepo
}

// NewProcessingService creates a new instance of the Processing Service
func NewProcessingService(data interfaces.DataRepo) processing {
	return processing{
		data: data,
	}
}

// User gets the User with the given ID
func (p *processing) User(id int) (types.User, error) {
	users := p.data.Users()
	return users.WithID(id)
}

// UserActionCount gets the count of all actions the user with the given ID has performed
func (p *processing) UserActionCount(id int) (int, error) {
	//start by checking if user exists, if they do not, throw an error
	//could possibly remove if intended behaviour is just to return 0
	_, err := p.User(id)
	if err != nil {
		return 0, err
	}

	//get all actions
	actions := p.data.Actions()
	//filter to just actions done by the user specified
	actionsByUser := actions.ByUserWithID(id)
	//return length of this array
	return len(actionsByUser), nil
}

// NextActions takes an action type and returns how often each other action occured after the given action as a probability
// if an action is omitted from the map it has not occured after the given action
func (p *processing) NextActions(actionType string) (map[string]float64, error) {
	actions := p.data.Actions() //get all actions

	//create a map of userIDs to a slice of actions
	//this is just to split up actions into groups performed by each user
	usersActions := make(map[int]types.Actions)
	for _, action := range actions {
		usersActions[action.UserID] = append(usersActions[action.UserID], action)
	}

	//given an action type, get the count of each next action
	counts := nextActionCounts(actionType, usersActions)

	//find the total number of actions that followed the given action type
	totalCount := 0
	for _, count := range counts {
		totalCount += count
	}

	//calculate probabilities (rounded to 2dp) of each next action given an action type
	probabilities := make(map[string]float64, len(counts))
	for actionType, count := range counts {
		probability := float64(count) / float64(totalCount)
		probabilities[actionType] = math.Round(probability*100) / 100
	}

	//this never returns an error so could just omit that from the returns
	//leaving in as intended behaviour may be to respond with a 404 if action does not exist, or if it is always a final action e.g. "DELETE_ACCOUNT"
	return probabilities, nil
}

// nextActionCounts takes an action type and a map of actions by userID
// it returns a count of how many times each other action occured after the given action
// if an action is omitted from the map it has not occured after the given action
func nextActionCounts(actionType string, usersActions map[int]types.Actions) map[string]int {
	counts := make(map[string]int) //create map to save counts into

	// for each user
	for _, actions := range usersActions {
		slices.SortStableFunc(actions, compareActionsCreatedAt) //ensure slice is sorted

		//then loop through all their actions
		for i, action := range actions {
			if i >= len(actions)-1 { //don't need to check the last action they've taken as there is no "next"
				break
			}
			if action.Type != actionType { //if they ever performed the given action
				continue
			}
			nextActionType := actions[i+1].Type                 //check what the next action they did was
			counts[nextActionType] = counts[nextActionType] + 1 //and count that
		}
	}
	//counts is a map where they keys are each action type of the next actions after the given action
	//and the values are the count of how many times each one occured
	return counts
}

// compareActionsCreatedAt returns -1 if a was created before b, 1 if b was created before a and 0 otherwise
func compareActionsCreatedAt(a, b types.Action) int {
	if a.CreatedAt.Before(b.CreatedAt) {
		return -1
	}
	if b.CreatedAt.Before(a.CreatedAt) {
		return 1
	}
	return 0
}
