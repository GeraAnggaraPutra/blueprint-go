package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	jwt "github.com/GeraAnggaraPutra/blueprint-go/handler/jwt"
	"github.com/GeraAnggaraPutra/blueprint-go/helpers/crypt"
	"github.com/GeraAnggaraPutra/blueprint-go/module"
	payload "github.com/GeraAnggaraPutra/blueprint-go/payload/auth"
	"github.com/GeraAnggaraPutra/blueprint-go/service"
)

type AuthController struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{service}
}

func (c *AuthController) LoginController(ctx echo.Context) error {
	var request payload.LoginRequest

	if err := ctx.Bind(&request); err != nil {
		return module.ResponseData(ctx, http.StatusBadRequest, nil, err.Error(), "Failed parse login request")
	}

	user, err := c.service.GetUserByEmailSvc(request.Email)
	if err != nil {
		return module.ResponseData(ctx, http.StatusNotFound, nil, err.Error(), "Email not found")
	}

	err = crypt.ComparePassword(user.Password, request.Password)
	if err != nil {
		return module.ResponseData(ctx, http.StatusBadRequest, nil, err.Error(), "Password didn't match")
	}

	payloadJWT := jwt.AccessTokenPayload{
		ID: user.ID,
	}

	accessToken, err := jwt.CreateJWT(payloadJWT)
	if err != nil {
		return module.ResponseData(ctx, http.StatusBadRequest, nil, err.Error(), "Failed create session")
	}

	userRes := payload.ToReadUserResponse(user)

	res := payload.LoginResponse{
		AccessToken: accessToken.Token,
		ExpiredAt:   accessToken.ExpiredAt,
		User:        userRes,
	}

	return module.ResponseData(ctx, http.StatusOK, res, nil, "Successfully login")
}
