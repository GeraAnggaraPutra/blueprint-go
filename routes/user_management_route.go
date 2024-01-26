package routes

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	controller "github.com/GeraAnggaraPutra/blueprint-go/controller/user_management"
	"github.com/GeraAnggaraPutra/blueprint-go/middleware"
	"github.com/GeraAnggaraPutra/blueprint-go/repository/user_management/query"
	service "github.com/GeraAnggaraPutra/blueprint-go/service/user_management"
)

func addUserManagementRoute(e *echo.Echo, db *sqlx.DB) {
	userQuery := query.NewUserQuery(db)
	userService := service.NewUserService(userQuery)
	userController := controller.NewUserController(userService)
	mddw := middleware.NewMiddlewareToken(db)

	userRoute := e.Group("/user")
	userRoute.Use(mddw.ValidateToken)

	userRoute.GET("", userController.ReadListUserCtrl, middleware.ValidateSuperAdmin)
}
