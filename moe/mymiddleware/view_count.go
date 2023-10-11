package mymiddleware

import (
	"SMOE/moe/database"
	"github.com/labstack/echo/v4"
	"log/slog"
)

// ViewCount todo 怎样别统计到爬虫
func ViewCount(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		go func() {
			if err := database.UpdateView(c.Param("cid")); err != nil {
				slog.Error(err.Error())
			}
		}()
		return next(c)
	}
}
