package handler

import (
	"github.com/labstack/echo/v4"
)

func FrontErr(err error, c echo.Context) {
	c.Render(400, "404.template", err)
	return
}

func BackErr(err error, c echo.Context) error {
	return c.Render(404, "admin-err.template", err)
}
