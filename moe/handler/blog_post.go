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
	// 递归查询 https://www.sqlite.org/lang_with.html
	if err := database.DB.Select(&qpu.Comments, `
		WITH RECURSIVE cte AS (
		SELECT * FROM smoe_comments WHERE parent=0 AND cid=? AND status=?
		UNION ALL
		SELECT s.* FROM smoe_comments AS s, cte AS c
		WHERE s.parent = c.coid
		ORDER BY ROWID DESC--深度优先
		)
		SELECT * FROM cte;`, cid, "approved"); err != nil {
		return err
	}
	return c.Render(200, "post.template", qpu)
}
