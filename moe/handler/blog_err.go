package handler

import (
	"github.com/labstack/echo/v4"
)

func FrontErr(err error, c echo.Context) {

	_ = c.Render(404, "404.template", err)
}
