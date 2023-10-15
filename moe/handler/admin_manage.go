package handler

import (
	"SMOE/moe/database"
	"github.com/labstack/echo/v4"
)

func ManagePost(c echo.Context) error {
	qpu := database.NewQPU()
	defer qpu.Free()
	req := &struct {
		CommStatus string `query:"commstatus" default:"approved" `
		Status     string `query:"status" default:"publish" `
		Page       int    `query:"page" default:"1"`
		Cid        int    `query:"cid" default:"1"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := qpu.GetPosts(req.Status, 10, req.Page); err != nil {
		return err
	}
	return c.Render(200, "manage-posts.template", qpu)
}

func ManagePage(c echo.Context) error {
	qpu := database.NewQPU()
	defer qpu.Free()
	req := &struct {
		CommStatus string `query:"commstatus" default:"approved" `
		Status     string `query:"status" default:"publish" `
		Page       int    `query:"page" default:"1"`
		Cid        int    `query:"cid" default:"1"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := qpu.GetPages(); err != nil {
		return err
	}
	return c.Render(200, "manage-pages.template", qpu)
}

func ManageComment(c echo.Context) error {
	qpu := database.NewQPU()
	defer qpu.Free()
	req := &struct {
		CommStatus string `query:"commstatus" default:"approved" `
		Status     string `query:"status" default:"publish" `
		Page       int    `query:"page" default:"1"`
		Cid        int    `query:"cid" default:"1"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	//if err := qpu.CommentsWithCid(); err != nil {
	//	return err
	//}
	//qpu.
	return c.Render(200, "manage-comments.template", qpu)
}

func ManageMedia(c echo.Context) error {
	qpu := database.NewQPU()
	defer qpu.Free()
	req := &struct {
		CommStatus string `query:"commstatus" default:"approved" `
		Status     string `query:"status" default:"publish" `
		Page       int    `query:"page" default:"1"`
		Cid        int    `query:"cid" default:"1"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	return c.Render(200, "manage-medias.template", qpu)
}
