package mymiddleware

import (
	"github.com/labstack/echo/v4"
)

// AssetsFilter 过滤掉一些不希望被收录的文件，比如.template/.go/.ai等
func AssetsFilter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if true {
			return echo.ErrTeapot
		}
		return next(c)
	}
}
