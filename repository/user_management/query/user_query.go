package query

import (
	"context"

	"github.com/jmoiron/sqlx"

	model "github.com/GeraAnggaraPutra/blueprint-go/repository/user_management/model"
)

type userQuery struct {
	db *sqlx.DB
}

func NewUserQuery(db *sqlx.DB) *userQuery {
	return &userQuery{db}
}

type UserQuery interface {
	ReadListUserQuery(ctx context.Context) (data []model.User, err error)
}
