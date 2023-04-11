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
	PageNum, _ := isNum(c.Param("num"))
	data := struct {
		PageArr    []query.Contents
		PostArr    []query.Contents
		PageNum    int
		MaxPageNum int
	}{
		query.PageArr(db),
		query.PostArr(db, "publish", 5, PageNum),
		PageNum,
		query.Count(db, "post", "publish"),
	}
	return c.Render(http.StatusOK, "index.template", data)
}

func BlogIndexAjax(c echo.Context) error {
	//pageNum, _ := Smoe.IsNum(c.Param("num"))
	//indexData := new(IndexData)
	////queryPost(&indexData.IndexPost, "publish", pageNum, 5)
	//indexData.PageNext = pageNum + 1
	return c.Render(http.StatusOK, "index-primary_ajax.template", nil)
}
