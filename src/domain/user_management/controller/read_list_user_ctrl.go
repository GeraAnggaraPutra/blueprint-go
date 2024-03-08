package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	payload "github.com/GeraAnggaraPutra/blueprint-go/src/domain/user_management/payload"
	"github.com/GeraAnggaraPutra/blueprint-go/src/module"
)

func (c *UserController) ReadListUserCtrl(ctx echo.Context) error {
	data, err := c.service.ReadListUserSvc(ctx.Request().Context())
	if err != nil {
		return module.ResponseData(ctx, http.StatusInternalServerError, nil, err.Error(), "failed read list user")
	}

	res := payload.ToReadUserResponses(data)

	return module.ResponseData(ctx, http.StatusOK, res, nil, "successfully read list user")
}
