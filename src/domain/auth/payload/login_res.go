package payload

import (
	"time"

	"github.com/GeraAnggaraPutra/blueprint-go/src/domain/auth/repository/model"
)

type UserResponse struct {
	ID        int64        `json:"id"`
	GUID      string       `json:"guid"`
	Email     string       `json:"email"`
	Role      RoleResponse `json:"role"`
	CreatedAt time.Time    `json:"created_at"`
	CreatedBy string       `json:"created_by"`
	UpdatedAt *time.Time   `json:"updated_at"`
	UpdatedBy *string      `json:"updated_by"`
}

type RoleResponse struct {
	GUID string `json:"guid"`
	Name string `json:"name"`
}

func ToReadUserResponse(entity model.User) (res UserResponse) {
	res.ID = entity.ID
	res.GUID = entity.GUID
	res.Email = entity.Email
	res.CreatedAt = entity.CreatedAt
	res.CreatedBy = entity.CreatedBy
	res.Role.GUID = entity.RoleGUID
	res.Role.Name = entity.RoleName

	if entity.UpdatedAt.Valid {
		res.UpdatedAt = &entity.UpdatedAt.Time
	}

	if entity.UpdatedBy.Valid {
		res.UpdatedBy = &entity.UpdatedBy.String
	}

	return
}

type LoginResponse struct {
	GUID        string       `json:"guid"`
	AccessToken string       `json:"access_token"`
	ExpiredAt   time.Time    `json:"expired_at"`
	CreatedAt   time.Time    `json:"created_at"`
	User        UserResponse `json:"user"`
}
