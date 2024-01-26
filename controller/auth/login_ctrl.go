package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/GeraAnggaraPutra/blueprint-go/handler/jwt"
	"github.com/GeraAnggaraPutra/blueprint-go/helpers/crypt"
	"github.com/GeraAnggaraPutra/blueprint-go/module"
	payload "github.com/GeraAnggaraPutra/blueprint-go/payload/auth"
	"github.com/GeraAnggaraPutra/blueprint-go/repository/auth/model"
)

func (c *AuthController) LoginCtrl(ctx echo.Context) error {
	var request payload.LoginRequest

	if err := ctx.Bind(&request); err != nil {
		return module.ResponseData(ctx, http.StatusBadRequest, nil, err.Error(), "failed parse login request")
	}

	user, err := c.service.ReadUserByEmailSvc(ctx.Request().Context(), request.Email)
	if err != nil {
		return module.ResponseData(ctx, http.StatusNotFound, nil, err.Error(), "email not found")
	}

	err = crypt.ComparePassword(user.Password, request.Password)
	if err != nil {
		return module.ResponseData(ctx, http.StatusBadRequest, nil, err.Error(), "password didn't match")
	}

	uuid := crypt.GenerateUUID()

	payloadJWT := jwt.AccessTokenPayload{
		GUID:     uuid,
		UserGUID: user.GUID,
		Role:     user.RoleName,
	}

	accessToken, err := jwt.CreateJWT(payloadJWT)
	if err != nil {
		return module.ResponseData(ctx, http.StatusBadRequest, nil, err.Error(), "failed create session")
	}

	session, err := c.service.CreateSessionSvc(ctx.Request().Context(),
		model.Session{
			GUID:      uuid,
			UserGUID:  user.GUID,
			Token:     accessToken.Token,
			UserAgent: ctx.Request().UserAgent(),
			IPAddress: ctx.Request().RemoteAddr,
			ExpiredAt: accessToken.ExpiredAt,
			CreatedAt: time.Now(),
		},
	)

	userRes := payload.ToReadUserResponse(user)

	res := payload.LoginResponse{
		GUID:        session.GUID,
		AccessToken: accessToken.Token,
		ExpiredAt:   accessToken.ExpiredAt,
		CreatedAt:   session.CreatedAt,
		User:        userRes,
	}

	return module.ResponseData(ctx, http.StatusOK, res, nil, "successfully login")
}
