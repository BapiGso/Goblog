package admin

import (
	"github.com/BapiGso/SMOE/moe/query"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
)

func LogAccess(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	req := new(Param)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	data := struct {
		AccessArr []query.Access
	}{
		query.AccessArr(db, 10, req.Page),
	}
	return c.Render(200, "log-access.template", data)
}
