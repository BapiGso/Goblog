package post

import (
	"github.com/BapiGso/SMOE/moe/query"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Index(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	cid, _ := IsNum(c.Param("cid"))
	data := struct {
		Post     []query.Contents
		TestPost query.Contents
	}{
		query.QueryWithCid(db, cid),
		query.TestQueryPostWithCid(db, cid),
	}
	return c.Render(http.StatusOK, "testpost.template", data)
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
