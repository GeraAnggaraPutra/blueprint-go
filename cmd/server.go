package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/GeraAnggaraPutra/blueprint-go/config"
	"github.com/GeraAnggaraPutra/blueprint-go/src/api"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.NewDatabasePotgresql()
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	defer db.Close()

	e := echo.New()
	e.HideBanner = true

	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		log.Printf("Error: %v", err)
		e.DefaultHTTPErrorHandler(err, ctx)
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return c.Response().Status < http.StatusBadRequest
		},
	}))

	log.Printf("serving REST HTTP server | config: name=%s, host=%s, port=%s", os.Getenv("APP_NAME"), os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))

	err = api.Routes(e, db)
	if err != nil {
		log.Printf("Error start the server with err: %s", err)
	}
}
