package query

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/GeraAnggaraPutra/blueprint-go/repository/auth/dto"
	model "github.com/GeraAnggaraPutra/blueprint-go/repository/auth/model"
)

type authQuery struct {
	db *sqlx.DB
}

func NewAuthQuery(db *sqlx.DB) *authQuery {
	return &authQuery{db}
}

type AuthQuery interface {
	ReadUserByEmailQuery(ctx context.Context, email string) (data model.User, err error)
	CreateSessionQuery(ctx context.Context, arg dto.CreateSessionParams) (data model.Session, err error)
}
