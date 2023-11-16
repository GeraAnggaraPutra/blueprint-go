package auth

import (
	"github.com/labstack/echo/v4"

	"github.com/GeraAnggaraPutra/blueprint-go/middleware"
)

func ClaimToken(c echo.Context) (response middleware.JWTClaim) {
	user := c.Get("jwt-res")
	response = user.(middleware.JWTClaim)
	return
}
