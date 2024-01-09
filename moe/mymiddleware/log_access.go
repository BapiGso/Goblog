package mymiddleware

import (
	"SMOE/moe/database"
	"github.com/labstack/echo/v4"
	"log/slog"
)

// LogAccess todo 不要依赖database
func LogAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		go func() {
			data := map[string]any{
				"ua":    c.Request().UserAgent(),
				"url":   c.Request().RequestURI,
				"query": c.Request(),
				"ip":    c.RealIP(),

				"path": c.Path(),
			}
			err := database.InsertAccess(data)
			if err != nil {
				slog.Error(err.Error())
			}
		}()
		return next(c)
	}
}
