package service

import (
	"github.com/GeraAnggaraPutra/blueprint-go/model"
	"github.com/GeraAnggaraPutra/blueprint-go/repository"

)

type Service interface {
	GetUserSvc() (data []model.User, err error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) GetUserSvc() (data []model.User, err error) {
	data, err = s.repository.GetUserQuery()
	if err != nil {
		return
	}

	return
}
