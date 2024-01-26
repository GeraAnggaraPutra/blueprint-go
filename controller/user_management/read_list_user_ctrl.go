package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/GeraAnggaraPutra/blueprint-go/module"
	payload "github.com/GeraAnggaraPutra/blueprint-go/payload/user_management"
)

func (c *UserController) ReadListUserCtrl(ctx echo.Context) error {
	data, err := c.service.ReadListUserSvc(ctx.Request().Context())
	if err != nil {
		return module.ResponseData(ctx, http.StatusInternalServerError, nil, err.Error(), "failed read list user")
	}

	res := payload.ToReadUserResponses(data)

	return module.ResponseData(ctx, http.StatusOK, res, nil, "successfully read list user")
}
