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
	qpu := database.NewQPU()
	defer qpu.Free()
	PageNum, err := validateNum(c.Param("num"))
	if err != nil {
		return err
	}
	err = qpu.GetPages()
	if err != nil {
		return err
	}
	err = qpu.GetPosts("publish", 5, PageNum)
	if err != nil {
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
