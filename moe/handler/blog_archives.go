package handler

import (
	"SMOE/moe/database"
	"github.com/labstack/echo/v4"
)

func Archives(c echo.Context) error {
	qpu := database.NewQPU()
	defer qpu.Free()
	//todo limit
	if err := qpu.GetPosts("publish", 100, 0); err != nil {
		return err
	}
	return c.Render(200, "page-timeline.template", qpu)
}
