package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var (
	Hs256 = jwt.SigningMethodHS256
)

func GenerateToken(method jwt.SigningMethod, payload jwt.Claims, signKey string) (string, error) {
	t := jwt.NewWithClaims(method, payload)
	return t.SignedString([]byte(signKey))
}

// / DecodeToken decodes the JWT token and fills the passed claims object of type T with the parsed claims.
// T must implement the jwt.Claims interface.
func DecodeToken[T jwt.Claims](tokenString string, method jwt.SigningMethod, signKey string, claims *T) error {
	// Ensure that the passed claims object satisfies the jwt.Claims interface.
	claimsInterface, ok := interface{}(claims).(jwt.Claims)
	if !ok {
		return fmt.Errorf("invalid claims type: %T does not implement jwt.Claims", claims)
	}

	// Parse the token with the provided claims type
	token, err := jwt.ParseWithClaims(tokenString, claimsInterface, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is as expected
		if token.Method != method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signKey), nil
	})

	// If there was an error during parsing, return it
	if err != nil {
		return err
	}

	// If the token is invalid, return an error
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	// Claims have been populated into passed *claims, no need to return anything else
	return nil
}
