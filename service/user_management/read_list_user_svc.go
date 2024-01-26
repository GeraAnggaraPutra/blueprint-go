package service

import (
	"context"
	"log"

	"github.com/GeraAnggaraPutra/blueprint-go/repository/user_management/model"
)

func (s *userService) ReadListUserSvc(
	ctx context.Context,
) (data []model.User, err error) {
	data, err = s.userQuery.ReadListUserQuery(ctx)
	if err != nil {
		log.Printf("failed read list user with err: %s", err)
		return
	}

	return
}
