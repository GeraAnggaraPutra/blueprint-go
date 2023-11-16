package payload

type LoginRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}
