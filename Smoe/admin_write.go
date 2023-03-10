package smoe

import (
	"SMOE/smoe/query"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Smoe) WritePost(c echo.Context) error {
	req := new(Param)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	//if req.Cid == 0 {
	//	return c.Render(200, "write-post.template", nil)
	//}
	data := struct {
		Post []query.Contents
	}{
		query.QueryWithCid(s.Db, req.Cid),
	}
	return c.Render(200, "write-post.template", data)

}

func (s *Smoe) WritePage(c echo.Context) error {
	req := new(Param)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	//if req.Cid == 0 {
	//	return c.Render(200, "write-page.template", nil)
	//}
	data := struct {
		Page []query.Contents
	}{
		query.QueryWithCid(s.Db, req.Cid),
	}
	return c.Render(200, "write-page.template", data)
}
