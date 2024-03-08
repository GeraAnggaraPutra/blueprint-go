package service

import (
	"context"
	"log"

	"github.com/GeraAnggaraPutra/blueprint-go/src/domain/auth/repository/model"
)

func (s *authService) ReadUserByEmailSvc(
	ctx context.Context,
	email string,
) (data model.User, err error) {
	data, err = s.authQuery.ReadUserByEmailQuery(ctx, email)
	if err != nil {
		log.Printf("failed read user by email with err: %s", err)
		return
	}

	return
}
