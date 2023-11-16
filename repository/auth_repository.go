package repository

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/GeraAnggaraPutra/blueprint-go/model"
)

type AuthRepository interface {
	GetUserByEmailQuery(email string) (data model.User, err error)
}

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) GetUserByEmailQuery(email string) (data model.User, err error) {
	var db = r.db

	const stmt = `
		SELECT 
			id, 
			email,
			password, 
			created_at, 
			updated_at
		FROM 
			users
		WHERE 
			email = $1
	`

	err = db.QueryRow(stmt, email).Scan(
		&data.ID,
		&data.Email,
		&data.Password,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error get data user by email with err: %s", err)
		return
	}

	return
}
