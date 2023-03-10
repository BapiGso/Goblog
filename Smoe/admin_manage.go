package smoe

import (
	"SMOE/smoe/query"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Param 类名首字母一定要大写，又被坑一次
type Param struct {
	CommStatus string `query:"status" default:"approved" `
	Status     string `query:"status" default:"publish" `
	Page       uint64 `query:"page" default:"1"`
	Cid        uint64 `query:"cid" default:"1"`
}

func (s *Smoe) ManagePost(c echo.Context) error {
	req := new(Param)

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	data := struct {
		PostArr []query.Contents
	}{
		query.QueryPostArr(s.Db, req.Status, 20, req.Page),
	}

	return c.Render(200, "manage-posts.template", data)
}

func (s *Smoe) ManagePage(c echo.Context) error {
	data := struct {
		PageArr []query.Contents
	}{
		query.QueryPageArr(s.Db),
	}

	return c.Render(200, "manage-pages.template", data)
}

func (s *Smoe) ManageComment(c echo.Context) error {
	req := new(Param)

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	data := struct {
		CommArr []query.Comments
	}{
		query.QueryCommentsArr(s.Db, req.CommStatus, 20, req.Page),
	}
	return c.Render(200, "manage-comments.template", data)
}

func (s *Smoe) ManageMedia(c echo.Context) error {
	req := new(Param)

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	data := struct {
		MediaArr []query.Contents
	}{
		query.QueryMedia(s.Db, 20, req.Page),
	}
	return c.Render(200, "manage-medias.template", data)
}
