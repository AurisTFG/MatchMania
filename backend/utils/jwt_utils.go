package utils

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GetJwtClaim[T any](tokenStr string, claimName string) (T, error) {
	var zero T

	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return zero, fmt.Errorf("parsing JWT: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return zero, errors.New("invalid JWT claims")
	}

	untypedClaim, exists := claims[claimName]
	if !exists {
		return zero, fmt.Errorf("claim '%s' not found", claimName)
	}

	typedClaim, ok := untypedClaim.(T)
	if !ok {
		return zero, fmt.Errorf("claim '%s' is not of expected type", claimName)
	}

	return typedClaim, nil
}
