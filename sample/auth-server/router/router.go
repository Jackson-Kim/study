package router

import (
	"auth-server/cookie"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) error {
	e.POST("/cookie/login", cookie.Login)
	// e.GET("/jwt/login", jwt.AuthJWT)
	// e.GET("/atrt/login", atrt.AuthATRT)
	// e.GET("/oauth/login", oauth.AuthOAuth)
	return nil
}
