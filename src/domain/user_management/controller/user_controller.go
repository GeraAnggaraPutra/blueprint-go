package controller

import service "github.com/GeraAnggaraPutra/blueprint-go/src/domain/user_management/service"

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service}
}
