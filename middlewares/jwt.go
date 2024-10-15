package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
  config := middleware.JWTConfig{
    SigningKey: []byte("secret"),
  }
  return middleware.JWTWithConfig(config)
}
