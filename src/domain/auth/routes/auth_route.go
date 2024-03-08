package routes

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	controller "github.com/GeraAnggaraPutra/blueprint-go/src/domain/auth/controller"
	"github.com/GeraAnggaraPutra/blueprint-go/src/domain/auth/repository/query"
	service "github.com/GeraAnggaraPutra/blueprint-go/src/domain/auth/service"
)

func AddAuthRoute(e *echo.Echo, db *sqlx.DB) {
	authQuery := query.NewAuthQuery(db)
	authService := service.NewAuthService(authQuery)
	authController := controller.NewAuthController(authService)

	authRoute := e.Group("/auth")
	authRoute.POST("/login", authController.LoginCtrl)
}
