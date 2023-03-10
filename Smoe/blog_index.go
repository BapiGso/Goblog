package smoe

import (
	"github.com/BapiGso/SMOE/smoe/query"
	"github.com/labstack/echo/v4"
	"net/http"
)

// FrontIndex TODO 加载更多、ajax
func (s *Smoe) BlogIndex(c echo.Context) error {
	//判断页数查数据库
	PageNum, _ := IsNum(c.Param("num"))
	data := struct {
		PageArr    []query.Contents
		PostArr    []query.Contents
		PageNum    uint64
		MaxPageNum uint64
	}{
		query.QueryPageArr(s.Db),
		query.QueryPostArr(s.Db, "publish", 5, PageNum),
		PageNum,
		query.QueryCount(s.Db, "post", "publish"),
	}
	return c.Render(http.StatusOK, "index.template", data)
}

func (s *Smoe) BlogIndexAjax(c echo.Context) error {
	//pageNum, _ := Smoe.IsNum(c.Param("num"))
	//indexData := new(IndexData)
	////queryPost(&indexData.IndexPost, "publish", pageNum, 5)
	//indexData.PageNext = pageNum + 1
	return c.Render(http.StatusOK, "index-primary_ajax.template", nil)
}
