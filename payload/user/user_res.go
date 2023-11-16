package payload

import (
	"time"

	"github.com/GeraAnggaraPutra/blueprint-go/model"
)

type readUserResponse struct {
	ID        int64      `json:"id"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func ToReadUserResponse(entity model.User) (response readUserResponse) {
	response.ID = entity.ID
	response.Email = entity.Email
	response.CreatedAt = entity.CreatedAt

	if entity.UpdatedAt.Valid {
		response.UpdatedAt = &entity.UpdatedAt.Time
	}

	return
}

func ToReadUserResponses(entities []model.User) (response []*readUserResponse) {
	response = make([]*readUserResponse, len(entities))

	for i := range entities {
		response[i] = new(readUserResponse)
		data := ToReadUserResponse(entities[i])
		response[i] = &data
	}

	return
}
