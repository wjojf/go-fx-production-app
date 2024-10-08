package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
)

func GenerateToken(method jwt.SigningMethod, payload jwt.Claims, signKey string) (string, error) {
	t := jwt.NewWithClaims(method, payload)
	return t.SignedString([]byte(signKey))
}

// DecodeToken decodes the JWT token and returns the claims of type T.
func DecodeToken[T jwt.StandardClaims](tokenString string, method jwt.SigningMethod, signKey string) (T, error) {
	var empty T // Return an empty T in case of an error

	// Create a new instance of T by using reflect (since we cannot directly use new(T) due to interface constraints)
	claims := reflect.New(reflect.TypeOf(empty)).Interface().(jwt.Claims)

	// Parse the token with the provided claims type
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is as expected
		if token.Method != method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signKey), nil
	})

	// If there was an error during parsing, return it
	if err != nil {
		return empty, err
	}

	// If the token is valid and claims are correct
	if !token.Valid {
		return empty, nil
	}

	// Otherwise, return an error for invalid token
	claimsTyped, ok := claims.(T)
	if !ok {
		return empty, fmt.Errorf("invalid claims type: %T", claims)
	}

	return claimsTyped, nil
}
