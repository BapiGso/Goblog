package mymiddleware

import "github.com/labstack/echo/v4"

func CommentMail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().After(func() {
			if c.Response().Status == 200 {

			}
		})
		return next(c)
	}
}
