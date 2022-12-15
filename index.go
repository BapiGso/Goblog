package main

import (
	"github.com/labstack/echo/v4"
	"main/smoe"
	"net/http"
)

// TODO 加载更多、ajax
// TODO 文章载入动画
func Index(c echo.Context) error {
	//判断页数查数据库
	PageNum, _ := Smoe.IsNum(c.Param("num"))
	data := struct {
		PageArr    []Smoe.Contents
		PostArr    []Smoe.Contents
		PageNum    uint64
		MaxPageNum uint64
	}{
		s.QueryPageArr(),
		s.QueryPostArr("publish", 5, PageNum),
		PageNum,
		s.QueryCount("post", "publish"),
	}
	return c.Render(http.StatusOK, "index.template", data)
}

func IndexAjax(c echo.Context) error {
	//pageNum, _ := Smoe.IsNum(c.Param("num"))
	//indexData := new(IndexData)
	////queryPost(&indexData.IndexPost, "publish", pageNum, 5)
	//indexData.PageNext = pageNum + 1
	return c.Render(http.StatusOK, "index-primary_ajax.template", nil)
}
