package controller

import "github.com/GeraAnggaraPutra/blueprint-go/src/domain/auth/service"

type AuthController struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{service}
}
