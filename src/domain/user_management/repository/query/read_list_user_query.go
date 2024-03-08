package query

import (
	"context"

	model "github.com/GeraAnggaraPutra/blueprint-go/src/domain/user_management/repository/model"
)

func (r *userQuery) ReadListUserQuery(
	ctx context.Context,
) (data []model.User, err error) {
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
	`

	err = r.db.SelectContext(ctx, &data, stmt)
	if err != nil {
		return
	}

	return
}
