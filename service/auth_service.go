package service

import (
	"log"

	"github.com/GeraAnggaraPutra/blueprint-go/model"
	"github.com/GeraAnggaraPutra/blueprint-go/repository"
)

type AuthService interface {
	GetUserByEmailSvc(email string) (data model.User, err error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(repository repository.AuthRepository) *authService {
	return &authService{repository}
}

func (s *authService) GetUserByEmailSvc(email string) (data model.User, err error) {
	data, err = s.authRepository.GetUserByEmailQuery(email)
	if err != nil {
		log.Printf("Error get user by email with err: %s", err)
		return
	}

	return
}
