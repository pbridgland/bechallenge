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
