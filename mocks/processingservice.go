package mocks

import "bechallenge/types"

type ProcessingService struct {
	NextUserResult types.User
	NextUserErr    error
}

func (p *ProcessingService) User(id int) (types.User, error) {
	return p.NextUserResult, p.NextUserErr
}
