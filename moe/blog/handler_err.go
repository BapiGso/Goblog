package blog

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func FrontErr(err error, c echo.Context) {
	c.Render(http.StatusNotFound, "404.template", err)
}

func BackErr(err error, c echo.Context) {
	c.Render(http.StatusNotFound, "admin-err.template", err)
}
