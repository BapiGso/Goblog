package admin

import (
	"github.com/labstack/echo/v4"
)

func Setting(c echo.Context) error {
	return c.Render(200, "setting.template", nil)
}
