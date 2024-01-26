package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/GeraAnggaraPutra/blueprint-go/db"
)

func Init() error {
	e := echo.New()

	db, err := db.Init()
	if err != nil {
		return err
	}
	defer db.Close()

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

	addAuthRoute(e, db)
	addUserManagementRoute(e, db)

	return e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
