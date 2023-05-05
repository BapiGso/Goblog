package blog

import (
	"github.com/BapiGso/SMOE/moe/query"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
)

// FrontIndex TODO 加载更多、ajax
func BlogIndex(c echo.Context) error {
	//判断页数查数据库
	db := c.Get("db").(*sqlx.DB)
	PageNum, err := isNum(c.Param("num"))
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	data := struct {
		PageArr     []query.Contents
		PostArr     []query.Contents
		NextPageNum int
	}{
		query.PageArr(db),
		query.PostArr(db, "publish", 5, PageNum),
		query.NextPageNum(db, "publish", 5, PageNum),
	}
	return c.Render(http.StatusOK, "index.template", data)
}

func BlogIndexAjax(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	PageNum, err := isNum(c.Param("num"))
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	data := struct {
		PostArr     []query.Contents
		NextPageNum int
	}{
		query.PostArr(db, "publish", 5, PageNum),
		query.NextPageNum(db, "publish", 5, PageNum),
	}
	return c.Render(http.StatusOK, "index-primary_ajax.template", data)
}
