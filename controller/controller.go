package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/GeraAnggaraPutra/blueprint-go/service"
)

type Controller struct {
	service service.Service
}

func NewController(service service.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) HelloWorldController(ctx echo.Context) error {
	// Perform appropriate processing here
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Hello World",
	})
}
