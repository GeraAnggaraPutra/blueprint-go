package controller

import service "github.com/GeraAnggaraPutra/blueprint-go/service/user_management"

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service}
}
