package handler

import (
	"SMOE/moe/database"
	"github.com/labstack/echo/v4"
)

func Post(c echo.Context) error {
	qpu := new(database.QPU)
	cid, err := validateNum(c.Param("cid"))
	if err != nil {
		return err
	}
	if err := database.DB.Select(&qpu.Contents, `
		SELECT * FROM smoe_contents 
		WHERE cid=? AND status=?
		AND type='post'`, cid, "publish"); err != nil {
		return err
	}
	if len(qpu.Contents) == 0 {
		return echo.NotFoundHandler(c)
	}
	if err := database.DB.Select(&qpu.Comments, `SELECT * FROM  smoe_comments 
		WHERE cid=?
		AND status=?
		ORDER BY created`, cid, "approved"); err != nil {
		return err
	}
	return c.Render(200, "post.template", qpu)
}
