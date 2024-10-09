package jwt_test

import (
	"github.com/dgrijalva/jwt-go"
	lib "github.com/wjojf/go-uber-fx/pkg/jwt"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {

	token, err := lib.GenerateToken(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			Id:        "testId",
		},
		"secretJwtKey",
	)

	if err != nil {
		t.Error(err)
	}

	t.Logf("Token: %s", token)
}

func TestValidatedToken(t *testing.T) {

	token, err := lib.GenerateToken(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			Id:        "testId",
		},
		"secretJwtKey",
	)

	if err != nil {
		t.Error(err)
	}

	var claims = &jwt.StandardClaims{}
	err = lib.DecodeToken(token, jwt.SigningMethodHS256, "testKey", claims)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Claims: %v", claims)

}
