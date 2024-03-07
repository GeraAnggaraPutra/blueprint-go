package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type JWTClaim struct {
	GUID     string
	UserGUID string
	Role     string
	jwt.StandardClaims
}

type MiddlewareToken struct {
	db *sqlx.DB
}

func NewMiddlewareToken(db *sqlx.DB) *MiddlewareToken {
	return &MiddlewareToken{
		db: db,
	}
}

func (mt *MiddlewareToken) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqToken string

		headerDataToken := c.Request().Header.Get("Authorization")

		splitToken := strings.Split(headerDataToken, "Bearer ")
		if len(splitToken) > 1 {
			reqToken = splitToken[1]
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		jwtResponse, err := ClaimsJWT(reqToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "You are not login yet")
		}

		if !mt.checkSession(jwtResponse.GUID) {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		c.Set("claims", jwtResponse)

		return next(c)
	}
}

func ClaimsJWT(token string) (response JWTClaim, err error) {
	var (
		secretKey = []byte(os.Getenv("AUTH_ACCESS_TOKEN_SECRET_KEY"))
		jwtClaims JWTClaim
		jwtToken  *jwt.Token
	)

	jwtToken, err = jwt.ParseWithClaims(
		token,
		&jwtClaims,
		func(jwtToken *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)
	if err != nil {
		return
	}

	if jwtToken == nil || !jwtToken.Valid {
		return
	}

	response = JWTClaim{
		GUID:     jwtClaims.GUID,
		UserGUID: jwtClaims.UserGUID,
		Role:     jwtClaims.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwtClaims.ExpiresAt,
		},
	}

	return
}

func (mt *MiddlewareToken) checkSession(guid string) (exists bool) {
	const stmt = `
		SELECT EXISTS (
			SELECT 
				1 
			FROM 
				sessions 
			WHERE 
				guid = $1
		)
	`
	_ = mt.db.Get(&exists, stmt, guid)

	return
}

func ValidateSuperAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		claims, _ := c.Get("claims").(JWTClaim)

		if claims.Role != "Super Admin" {
			return echo.NewHTTPError(
				http.StatusForbidden,
				"your permission is not allowed to access this resource",
			)
		}

		return next(c)
	}
}
