package index

import (
	"github.com/BapiGso/SMOE/moe/query"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// FrontIndex TODO 加载更多、ajax
func BlogIndex(c echo.Context) error {
	//判断页数查数据库
	db := c.Get("db").(*sqlx.DB)
	PageNum, _ := IsNum(c.Param("num"))
	data := struct {
		PageArr    []query.Contents
		PostArr    []query.Contents
		PageNum    uint64
		MaxPageNum uint64
	}{
		query.QueryPageArr(db),
		query.QueryPostArr(db, "publish", 5, PageNum),
		PageNum,
		query.QueryCount(db, "post", "publish"),
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

// IsNum 首页返回1，不是数字返回err调用404，其他为对应页数
func IsNum(numstr string) (uint64, error) {
	if numstr == "" {
		return 1, nil
	}
	num, err := strconv.ParseUint(numstr, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}
