package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/GeraAnggaraPutra/blueprint-go/controller"
	"github.com/GeraAnggaraPutra/blueprint-go/db"
	"github.com/GeraAnggaraPutra/blueprint-go/repository"
	"github.com/GeraAnggaraPutra/blueprint-go/service"
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

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	// Routes
	e.GET("", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": fmt.Sprintf("%s API is Running", os.Getenv("APP_NAME")),
		})
	})

	e.GET("/hello-world", controller.HelloWorldController)
	e.GET("/users", controller.GetUserController)

	AddAuthRoute(e, db)

	return e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
