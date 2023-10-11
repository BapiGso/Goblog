package handler

import (
	"SMOE/moe/database"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Test(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	data := make([]database.Contents, 0, 1)
	_ = db.Select(&data, `SELECT * FROM typecho_contents WHERE cid=?`, 11)
	return c.JSON(200, data)
}
