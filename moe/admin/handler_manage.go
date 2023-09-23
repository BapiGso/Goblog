package admin

import (
	"SMOE/moe/database"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Param 类名首字母一定要大写，又被坑一次
type Param struct {
	CommStatus string `query:"commstatus" default:"approved" `
	Status     string `query:"status" default:"publish" `
	Page       int    `query:"page" default:"1"`
	Cid        int    `query:"cid" default:"1"`
}

func ManagePost(c echo.Context) error {
	_ = c.Get("db").(*sqlx.DB)
	req := new(Param)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	data := struct {
		PostArr []database.Contents
	}{
		//database.PostArr(db, req.Status, 10, req.Page),
	}
	return c.Render(200, "manage-posts.template", data)
}

func ManagePage(c echo.Context) error {
	_ = c.Get("db").(*sqlx.DB)
	data := struct {
		PageArr []database.Contents
	}{
		//database.PageArr(db),
	}

	return c.Render(200, "manage-pages.template", data)
}

func ManageComment(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	req := new(Param)

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	data := struct {
		CommArr []database.CommsWithTitleMix
	}{
		database.CommsWithTitle(db, req.CommStatus, 5, req.Page),
	}
	return c.Render(200, "manage-comments.template", data)
}

func ManageMedia(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	req := new(Param)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	data := struct {
		MediaArr []database.MediasWithTitleMix
	}{
		database.MediasWithTitle(db, 10, req.Page),
	}
	return c.Render(200, "manage-medias.template", data)
}
