package routes

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	controller "github.com/GeraAnggaraPutra/blueprint-go/controller/auth"
	"github.com/GeraAnggaraPutra/blueprint-go/repository/auth/query"
	service "github.com/GeraAnggaraPutra/blueprint-go/service/auth"
)

func addAuthRoute(e *echo.Echo, db *sqlx.DB) {
	authQuery := query.NewAuthQuery(db)
	authService := service.NewAuthService(authQuery)
	authController := controller.NewAuthController(authService)

	authRoute := e.Group("/auth")
	authRoute.POST("/login", authController.LoginCtrl)
}
