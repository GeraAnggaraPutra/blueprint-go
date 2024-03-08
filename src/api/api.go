package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	authRoute "github.com/GeraAnggaraPutra/blueprint-go/src/domain/auth/routes"
	userManRoute "github.com/GeraAnggaraPutra/blueprint-go/src/domain/user_management/routes"
)

func Routes(e *echo.Echo, db *sqlx.DB) error {
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodConnect, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderAccept, echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderXCSRFToken, echo.HeaderContentType, echo.HeaderContentLength},
	}))

	// Routes
	e.GET("", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": fmt.Sprintf("%s API is Running", os.Getenv("APP_NAME")),
		})
	})

	e.Static("/media", "src/public")

	authRoute.AddAuthRoute(e, db)
	userManRoute.AddUserManagementRoute(e, db)

	return e.Start(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
}
