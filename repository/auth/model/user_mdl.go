package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64          `db:"id"`
	GUID      string         `db:"guid"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	RoleGUID  string         `db:"role_guid"`
	RoleName  string         `db:"role_name"`
	CreatedAt time.Time      `db:"created_at"`
	CreatedBy string         `db:"created_by"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
	UpdatedBy sql.NullString `db:"updated_by"`
}
