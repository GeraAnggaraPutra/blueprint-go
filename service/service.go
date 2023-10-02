package service

import "github.com/GeraAnggaraPutra/blueprint-go/repository"

type Service interface {
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}
