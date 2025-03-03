package mocks

import "bechallenge/types"

type ProcessingService struct {
	NextUserResult            types.User
	NextUserErr               error
	NextUserActionCountResult int
	NextUserActionCountErr    error
	NextResultForNextActions  map[string]float64
	NextErrorForNextActions   error
}

func (p *ProcessingService) NextActions(actionType string) (map[string]float64, error) {
	return p.NextResultForNextActions, p.NextErrorForNextActions
}

func (p *ProcessingService) User(id int) (types.User, error) {
	return p.NextUserResult, p.NextUserErr
}

func (p *ProcessingService) UserActionCount(id int) (int, error) {
	return p.NextUserActionCountResult, p.NextUserActionCountErr
}
