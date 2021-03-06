package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/8luebottle/Wells-Far-Go/api"
)

// EnableCORS enables cross-origin resource sharing.
// It sets Allow-Origin, Allow-Method, Allow-Header, and Allow-Credentials
func EnableCORS(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			api.LOCAL,
		},
		AllowMethods: []string{
			echo.GET, echo.PUT, echo.POST, echo.DELETE,
		},
		AllowHeaders: []string{
			echo.HeaderAccessControlAllowCredentials, echo.HeaderAcceptEncoding,
			echo.HeaderAuthorization, echo.HeaderContentType,
			echo.HeaderOrigin, echo.HeaderXRequestedWith,
		},
		AllowCredentials: true,
	}))
}
