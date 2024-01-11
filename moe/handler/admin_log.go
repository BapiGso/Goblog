package handler

import (
	"SMOE/moe/database"
	"github.com/labstack/echo/v4"
)

func LogAccess(c echo.Context) error {
	qpu := new(database.QPU)
	req := &struct {
		CommStatus string `query:"commstatus" default:"approved" `
		Status     string `query:"status" default:"publish" `
		Page       int    `query:"page" default:"1"`
		Cid        int    `query:"cid" default:"1"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	return c.Render(200, "log-access.template", qpu)
}
