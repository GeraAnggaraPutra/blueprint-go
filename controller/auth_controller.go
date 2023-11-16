package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	jwt "github.com/GeraAnggaraPutra/blueprint-go/handler/jwt"
	"github.com/GeraAnggaraPutra/blueprint-go/helpers/crypt"
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
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed parse login request",
		})
	}

	user, err := c.service.GetUserByEmailSvc(request.Email)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Email not found",
		})
	}

	err = crypt.ComparePassword(user.Password, request.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Password didn't match",
		})
	}

	payloadJWT := jwt.AccessTokenPayload{
		ID: user.ID,
	}

	accessToken, err := jwt.CreateJWT(payloadJWT)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed create session",
		})
	}

	res := payload.LoginResponse{
		AccessToken: accessToken.Token,
		ExpiredAt:   accessToken.ExpiredAt,
		User: payload.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
		},
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully login",
		"data":    res,
	})
}
