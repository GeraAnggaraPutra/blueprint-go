package service

import (
	"context"

	"github.com/GeraAnggaraPutra/blueprint-go/src/domain/user_management/repository/model"
	"github.com/GeraAnggaraPutra/blueprint-go/src/domain/user_management/repository/query"
)

type userService struct {
	userQuery query.UserQuery
}

func NewUserService(query query.UserQuery) *userService {
	return &userService{query}
}

type UserService interface {
	ReadListUserSvc(ctx context.Context) (data []model.User, err error)
}
