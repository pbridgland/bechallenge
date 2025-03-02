package services

import (
	"bechallenge/interfaces"
	"bechallenge/types"
)

type processing struct {
	data interfaces.DataService
}

func NewProcessingService(data interfaces.DataService) processing {
	return processing{
		data: data,
	}
}

func (p *processing) User(id int) (types.User, error) {
	users := p.data.Users()
	return users.WithID(id)
}

func (p *processing) UserActionCount(id int) (int, error) {
	//start by checking if user exists, if they do not, throw an error
	//could possibly remove if intended behaviour is just to return 0
	users := p.data.Users()
	_, err := users.WithID(id)
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
