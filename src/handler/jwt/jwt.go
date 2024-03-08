package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type AccessTokenPayload struct {
	GUID     string `json:"guid"`
	UserGUID string `json:"user_guid"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type JWTPayload struct {
	Token     string
	ExpiredAt time.Time
}

func CreateJWT(req AccessTokenPayload) (response JWTPayload, err error) {
	var (
		jwtToken  *jwt.Token
		secretKey = []byte(os.Getenv("AUTH_ACCESS_TOKEN_SECRET_KEY"))
	)

	expiredDuration, err := time.ParseDuration(os.Getenv("AUTH_ACCESS_TOKEN_EXPIRES"))
	if err != nil {
		return
	}

	expiredAt := time.Now().Add(expiredDuration)

	jwtClaims := &AccessTokenPayload{
		GUID:     req.GUID,
		UserGUID: req.UserGUID,
		Role:     req.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
		},
	}

	jwtToken = jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	token, err := jwtToken.SignedString(secretKey)
	if err != nil {
		return
	}

	response = JWTPayload{
		Token:     token,
		ExpiredAt: expiredAt,
	}

	return
}
