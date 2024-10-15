package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func RoleMiddleware(role string) echo.MiddlewareFunc {
  return func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
      user := c.Get("user").(*jwt.Token)
      claims := user.Claims.(jwt.MapClaims)
      userRole := claims["role"].(string)

      if userRole != role {
        return echo.ErrForbidden
      }

      return next(c)
    }
  }
}
