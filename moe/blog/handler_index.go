package blog

import (
	"SMOE/moe/database"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Index TODO 加载更多、ajax
func Index(c echo.Context) error {
	//判断页数查数据库
	qpu := database.NewQPU()
	defer qpu.Free()
	//_, err := session.Get("smoeSession", c)
	//if err != nil {
	//	return err
	//}
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
	return c.Render(http.StatusOK, "index.template", qpu)
}

func IndexAjax(c echo.Context) error {
	_ = c.Get("db").(*sqlx.DB)
	_, err := validateNum(c.Param("num"))
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.Render(http.StatusOK, "index-primary_ajax.template", nil)
}
