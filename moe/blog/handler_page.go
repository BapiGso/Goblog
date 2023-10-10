package blog

import (
	"SMOE/moe/database"
	"github.com/labstack/echo/v4"
)

func Page(c echo.Context) error {
	qpu := database.NewQPU()
	err := qpu.GetPage(c.Param("page"))
	if err != nil {
		return err
	}
	return c.Render(200, "page.template", qpu)
}
