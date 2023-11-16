package payload

import (
	"time"

	"github.com/GeraAnggaraPutra/blueprint-go/model"
)

type UserResponse struct {
	ID        int64      `json:"id"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func ToReadUserResponse(entity model.User) (response UserResponse) {
	response.ID = entity.ID
	response.Email = entity.Email
	response.CreatedAt = entity.CreatedAt

	if entity.UpdatedAt.Valid {
		response.UpdatedAt = &entity.UpdatedAt.Time
	}

	return
}

type LoginResponse struct {
	AccessToken string       `json:"access_token"`
	ExpiredAt   time.Time    `json:"expired_at"`
	User        UserResponse `json:"user"`
}
