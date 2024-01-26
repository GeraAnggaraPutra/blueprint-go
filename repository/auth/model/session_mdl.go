package model

import "time"

type Session struct {
	GUID      string    `db:"guid"`
	UserGUID  string    `db:"user_guid"`
	IPAddress string    `db:"ip_address"`
	Token     string    `db:"token"`
	UserAgent string    `db:"user_agent"`
	ExpiredAt time.Time `db:"expired_at"`
	CreatedAt time.Time `db:"created_at"`
}
