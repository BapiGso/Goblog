package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mcuadros/go-defaults"
	"main/smoe"
	"net/http"
)

// Param 类名首字母一定要大写，又被坑一次
type Param struct {
	CommStatus string `query:"status" default:"approved" `
	Status     string `query:"status" default:"publish" `
	Page       uint64 `query:"page" default:"1"`
	Cid        uint64 `query:"cid" default:"1"`
}

func ManagePost(c echo.Context) error {
	req := new(Param)
	defaults.SetDefaults(req)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	data := struct {
		PostArr []Smoe.Contents
	}{
		s.QueryPostArr(req.Status, 20, req.Page),
	}

	return c.Render(200, "manage-posts.template", data)
}

func ManagePage(c echo.Context) error {
	data := struct {
		PageArr []Smoe.Contents
	}{
		s.QueryPageArr(),
	}

	return c.Render(200, "manage-pages.template", data)
}

func ManageComment(c echo.Context) error {
	req := new(Param)
	defaults.SetDefaults(req)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	data := struct {
		CommArr []Smoe.Comments
	}{
		s.QueryCommentsArr(req.CommStatus, 20, req.Page),
	}
	return c.Render(200, "manage-comments.template", data)
}

func ManageMedia(c echo.Context) error {
	req := new(Param)
	defaults.SetDefaults(req)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	data := struct {
		MediaArr []Smoe.Contents
	}{
		s.QueryMedia(20, req.Page),
	}
	return c.Render(200, "manage-medias.template", data)
}
