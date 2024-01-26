package service

import (
	"context"

	"github.com/GeraAnggaraPutra/blueprint-go/repository/auth/model"
	"github.com/GeraAnggaraPutra/blueprint-go/repository/auth/query"
)

type authService struct {
	authQuery query.AuthQuery
}

func NewAuthService(query query.AuthQuery) *authService {
	return &authService{query}
}

type AuthService interface {
	ReadUserByEmailSvc(ctx context.Context, email string) (data model.User, err error)
	CreateSessionSvc(ctx context.Context, req model.Session) (data model.Session, err error)
}
