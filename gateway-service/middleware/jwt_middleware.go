package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid authorization header")
		}

		tokenRequest := strings.TrimPrefix(authHeader, "Bearer ")
		// Use JWT_SECRET (set in env_file) as the signing secret
		secretToken := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenRequest, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid signing method")
			}
			return []byte(secretToken), nil
		}, jwt.WithValidMethods([]string{"HS256"}))

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}

		// Set custom headers from JWT claims if available
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if v, ok := claims["id"].(string); ok {
				c.Request().Header.Set("TS-USER-ID", v)
			}
			if v, ok := claims["name"].(string); ok {
				c.Request().Header.Set("TS-USER-NAME", v)
			}
			if v, ok := claims["role"].(string); ok {
				c.Request().Header.Set("TS-USER-ROLE", v)
			}
			if v, ok := claims["email"].(string); ok {
				c.Request().Header.Set("TS-USER-EMAIL", v)
			}
		}
		return next(c)
	}
}
