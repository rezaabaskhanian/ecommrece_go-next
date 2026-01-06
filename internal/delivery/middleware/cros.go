package middleware

import "github.com/labstack/echo/v4"

func CORSMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			// پاسخ به preflight
			if c.Request().Method == "OPTIONS" {
				return c.NoContent(200)
			}

			return next(c)
		}
	}
}
