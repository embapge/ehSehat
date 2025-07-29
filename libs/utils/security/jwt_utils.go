package security

import (
	"context"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const jwtContextKey ctxKey = "jwt_payload"

// ExtractJWTFromContext returns JWT payload from context.
func ExtractJWTFromContext(ctx context.Context) (*JWTPayload, error) {
	payload, ok := ctx.Value(jwtContextKey).(*JWTPayload)
	if !ok || payload == nil {
		return nil, errors.New("JWT payload not found in context")
	}
	return payload, nil
}

// InjectJWTToContext parses the token and injects the payload into context.
func InjectJWTToContext(ctx context.Context, tokenString string, secret string) (context.Context, error) {
	payload, err := parseToken(tokenString, secret)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, jwtContextKey, payload), nil
}

// parseToken parses the JWT token string to a JWTPayload.
func parseToken(tokenString string, secret string) (*JWTPayload, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid JWT token")
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid JWT claims")
	}

	return &JWTPayload{
		ID:    getClaimString(claims, "id"),
		Name:  getClaimString(claims, "name"),
		Email: getClaimString(claims, "email"),
		Role:  getClaimString(claims, "role"),
	}, nil
}

func getClaimString(claims *jwt.MapClaims, key string) string {
	if v, ok := (*claims)[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
