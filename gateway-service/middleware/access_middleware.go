package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// AccessRole middleware: hanya mengizinkan role tertentu
func AccessRole(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role := c.Request().Header.Get("TS-USER-ROLE")
			for _, allowed := range allowedRoles {
				if strings.EqualFold(role, allowed) {
					return next(c)
				}
			}
			return echo.NewHTTPError(http.StatusForbidden, "Access denied: role not allowed")
		}
	}
}
