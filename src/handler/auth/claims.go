package auth

import (
	"github.com/labstack/echo/v4"

	"github.com/GeraAnggaraPutra/blueprint-go/src/middleware"
)

func Claims(c echo.Context) (response middleware.JWTClaim) {
	user := c.Get("jwt-res")

	response = user.(middleware.JWTClaim)

	return
}
