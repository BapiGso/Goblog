package admin

import (
	"SMOE/moe/database"
	"fmt"
	"github.com/labstack/echo/v4"

	"github.com/jmoiron/sqlx"
	"net/http"
)

func WritePost(c echo.Context) error {
	//db := c.Get("db").(*sqlx.DB)
	req := new(Param)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	//if req.Cid == 0 {
	//	return c.Render(200, "write-post.template", nil)
	//}
	data := struct {
		Post database.Contents
	}{
		//database.PostWithCid(db, req.Cid),
	}
	fmt.Println(req, data)
	return c.Render(200, "write-post.template", data)

}

func WritePage(c echo.Context) error {
	_ = c.Get("db").(*sqlx.DB)
	req := new(Param)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	data := struct {
		Page database.Contents
	}{
		//database.PageWithCid(db, req.Cid),
	}
	return c.Render(200, "write-page.template", data)
}
