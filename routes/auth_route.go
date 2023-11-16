package routes

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/GeraAnggaraPutra/blueprint-go/controller"
	"github.com/GeraAnggaraPutra/blueprint-go/repository"
	"github.com/GeraAnggaraPutra/blueprint-go/service"
)

func AddAuthRoute(e *echo.Echo, db *sqlx.DB) {
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	authRoute := e.Group("/auth")
	authRoute.POST("/login", authController.LoginController)
}
