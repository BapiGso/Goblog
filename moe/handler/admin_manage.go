package handler

import (
	"SMOE/moe/database"
	"github.com/labstack/echo/v4"
)

func Manage(c echo.Context) error {
	qpu := database.NewQPU()
	defer qpu.Free()
	req := &struct {
		Type       string `param:"type"       default:"post" `
		CommStatus string `query:"commstatus" default:"approved" `
		Status     string `query:"status"     default:"publish" `
		Page       int    `query:"page"       default:"1"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	switch req.Type {
	case "post":
		if err := qpu.GetPosts(req.Status, 10, req.Page); err != nil {
			return err
		}
	case "page":
		if err := qpu.GetPages(); err != nil {
			return err
		}
	case "comment":
		if err := qpu.GetComms(req.CommStatus, 10, req.Page); err != nil {
			//fmt.Println(len(qpu.CommArr))
			return err
		}
	}

	return c.Render(200, "manage.template", qpu)
}
