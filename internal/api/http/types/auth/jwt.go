package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtPayload struct {
	jwt.StandardClaims
	UserId string `json:"userId"`
}

func NewPayload(userId string, lifetime time.Duration) JwtPayload {
	return JwtPayload{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(lifetime).Unix(),
		},
		UserId: userId,
	}
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type TokenRefreshed struct {
	AccessToken string
	UserID      string
}
