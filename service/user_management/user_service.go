package service

import (
	"context"

	"github.com/GeraAnggaraPutra/blueprint-go/repository/user_management/model"
	"github.com/GeraAnggaraPutra/blueprint-go/repository/user_management/query"
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
