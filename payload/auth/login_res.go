package payload

import "time"

type UserResponse struct {
	ID        int64      `json:"id"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type LoginResponse struct {
	AccessToken string       `json:"access_token"`
	ExpiredAt   time.Time    `json:"expired_at"`
	User        UserResponse `json:"user"`
}
