package query

import (
	"context"
	"log"

	"github.com/GeraAnggaraPutra/blueprint-go/src/domain/auth/repository/dto"
	"github.com/GeraAnggaraPutra/blueprint-go/src/domain/auth/repository/model"
)

func (r *authQuery) CreateSessionQuery(
	ctx context.Context, arg dto.CreateSessionParams,
) (data model.Session, err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("failed begin tx with err: %s", err)
		return
	}

	defer func() {
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				log.Printf("error rollback: %v", errRollback)
				return
			}
		}
	}()

	const stmt = `
		INSERT INTO sessions
			(guid, user_guid, ip_address, token, user_agent, expired_at, created_at)
		VALUES
			($1, $2, $3, $4, $5, $6, (now() at time zone 'UTC')::TIMESTAMP)
		RETURNING
			guid, user_guid, ip_address, token, user_agent, expired_at, created_at
	`

	err = tx.QueryRowContext(ctx, stmt,
		arg.GUID,
		arg.UserGUID,
		arg.IPAddress,
		arg.Token,
		arg.UserAgent,
		arg.ExpiredAt,
	).Scan(
		&data.GUID,
		&data.UserGUID,
		&data.IPAddress,
		&data.Token,
		&data.UserAgent,
		&data.ExpiredAt,
		&data.CreatedAt,
	)
	if err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.Printf("error commit with err: %s", err)
		return
	}

	return
}
