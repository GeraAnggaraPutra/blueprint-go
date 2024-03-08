package query

import (
	"context"

	model "github.com/GeraAnggaraPutra/blueprint-go/src/domain/auth/repository/model"
)

func (r *authQuery) ReadUserByEmailQuery(
	ctx context.Context,
	email string,
) (data model.User, err error) {
	const stmt = `
		SELECT 
			u.id, 
			u.guid,
			u.email,
			u.password,
			u.role_guid,
			r.name AS role_name,
			u.created_at,
			u.created_by,
			u.updated_at,
			u.updated_by
		FROM 
			users u
		LEFT JOIN
			roles r
		ON u.role_guid = r.guid
		WHERE 
			email = $1
	`

	err = r.db.QueryRowContext(ctx, stmt, email).Scan(
		&data.ID,
		&data.GUID,
		&data.Email,
		&data.Password,
		&data.RoleGUID,
		&data.RoleName,
		&data.CreatedAt,
		&data.CreatedBy,
		&data.UpdatedAt,
		&data.UpdatedBy,
	)
	if err != nil {
		return
	}

	return
}
