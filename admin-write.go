package main

import (
	"github.com/labstack/echo/v4"
	Smoe "main/smoe"
	"net/http"
)

func WritePost(c echo.Context) error {
	req := new(Param)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	//if req.Cid == 0 {
	//	return c.Render(200, "write-post.template", nil)
	//}
	data := struct {
		Post []Smoe.Contents
	}{
		s.QueryWithCid(req.Cid),
	}
	return c.Render(200, "write-post.template", data)

}

func WritePage(c echo.Context) error {
	req := new(Param)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	//if req.Cid == 0 {
	//	return c.Render(200, "write-page.template", nil)
	//}
	data := struct {
		Page []Smoe.Contents
	}{
		s.QueryWithCid(req.Cid),
	}
	return c.Render(200, "write-page.template", data)
}
