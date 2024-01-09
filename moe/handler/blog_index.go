package handler

import (
	"SMOE/moe/database"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// Index TODO 加载更多、ajax
func Index(c echo.Context) error {
	//判断页数查数据库
	qpu := new(database.QPU)
	pageNum, err := validateNum(c.Param("num"))
	if err != nil {
		return err
	}
	//查询文章和独立页面
	if err := database.DB.Select(&qpu.Contents, `
		SELECT * FROM (
		  SELECT * FROM typecho_contents 
		  WHERE type='post' AND status= ?
		  ORDER BY ROWID DESC
		  LIMIT ? OFFSET ?  
		) t1
		UNION ALL
		SELECT * FROM  typecho_contents
		WHERE type='page'
		ORDER BY "order"
		`, "publish", 6, pageNum*5-5); err != nil {
		return err
	}
	if err := database.DB.Select(&qpu.Fields, `
		SELECT * FROM  typecho_fields
		WHERE name='coverList' AND cid IN (
			SELECT cid FROM typecho_contents
			WHERE type='post' AND status= ?
			ORDER BY ROWID DESC
			LIMIT ? OFFSET ?
		)
		`, "publish", 5, pageNum*5-5); err != nil {
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
	_, err := validateNum(c.Param("num"))
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.Render(200, "index-primary_ajax.template", nil)
}
