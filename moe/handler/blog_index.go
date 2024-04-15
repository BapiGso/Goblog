package handler

import (
	"SMOE/moe/database"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"strings"
)

// Index TODO 加载更多、ajax
func Index(c echo.Context) error {
	//判断页数查数据库
	qpu := new(database.QPU)
	req := &struct {
		PageNum int `param:"num" validate:"gte=0" default:"1"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	//查询文章和独立页面
	if err := database.DB.Select(&qpu.Contents, `
		SELECT * FROM (
		  SELECT * FROM smoe_contents 
		  WHERE type='post' AND status= ?
		  ORDER BY ROWID DESC
		  LIMIT ? OFFSET ?  
		)
		UNION ALL
		SELECT * FROM (
		SELECT * FROM  smoe_contents
		WHERE type='page'
		ORDER BY "created" DESC 
		)
		`, "publish", 6, req.PageNum*5-5); err != nil {
		return err
	}
	if !strings.Contains(c.Request().Header.Get(echo.HeaderAccept), echo.MIMETextHTML) {
		return c.Render(200, "index-primary_ajax.template", qpu)
	}
	return c.Render(200, "index.template", qpu)
}

// Deprecated: use x-request-with instead of
func IndexAjax(c echo.Context) error {
	_ = c.Get("db").(*sqlx.DB)
	return c.Render(200, "index-primary_ajax.template", nil)
}
