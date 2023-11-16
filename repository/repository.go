package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/GeraAnggaraPutra/blueprint-go/model"
)

type Repository interface {
	GetUserQuery() (data []model.User, err error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUserQuery() (data []model.User, err error) {
	var db = r.db

	query := `SELECT id, ..... FROM users`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User

		err := rows.Scan(
			&user.ID,
			&user.Email,
		)
		if err != nil {
			return []model.User{}, err
		}

		data = append(data, user)
	}

	return
}
