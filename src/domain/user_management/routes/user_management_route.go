package routes

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	controller "github.com/GeraAnggaraPutra/blueprint-go/src/domain/user_management/controller"
	"github.com/GeraAnggaraPutra/blueprint-go/src/domain/user_management/repository/query"
	service "github.com/GeraAnggaraPutra/blueprint-go/src/domain/user_management/service"
	"github.com/GeraAnggaraPutra/blueprint-go/src/middleware"
)

func AddUserManagementRoute(e *echo.Echo, db *sqlx.DB) {
	userQuery := query.NewUserQuery(db)
	userService := service.NewUserService(userQuery)
	userController := controller.NewUserController(userService)
	mddw := middleware.NewMiddlewareToken(db)

	userRoute := e.Group("/user")
	userRoute.Use(mddw.ValidateToken)

	userRoute.GET("", userController.ReadListUserCtrl, middleware.ValidateSuperAdmin)
}
