package controller

import service "github.com/GeraAnggaraPutra/blueprint-go/service/auth"

type AuthController struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{service}
}
