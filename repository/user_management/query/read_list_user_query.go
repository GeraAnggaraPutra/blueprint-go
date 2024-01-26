package query

import (
	"context"

	model "github.com/GeraAnggaraPutra/blueprint-go/repository/user_management/model"
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

	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u model.User
		if err = rows.Scan(
			&u.ID,
			&u.GUID,
			&u.Email,
			&u.Password,
			&u.RoleGUID,
			&u.RoleName,
			&u.CreatedAt,
			&u.CreatedBy,
			&u.UpdatedAt,
			&u.UpdatedBy,
		); err != nil {
			return
		}

		data = append(data, u)
	}

	return
}
