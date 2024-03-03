package mymiddleware

import (
	"SMOE/moe/database"
	"github.com/labstack/echo/v4"
	"log/slog"
	"time"
)

// LogAccess todo 不要依赖database
func LogAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().After(func() {
			_, err := database.DB.Exec(`INSERT INTO smoe_access_log (ua, url, ip, referer, time) VALUES (?, ?, ?, ?,?)`,
				c.Request().UserAgent(), c.Request().URL.String(), c.RealIP(), c.Request().Referer(), time.Now().Unix())
			if err != nil {
				//todo database is locked (5) (SQLITE_BUSY)
				slog.Error(err.Error())
			}
		})
		return next(c)
	}
}
