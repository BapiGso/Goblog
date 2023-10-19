package handler

import (
	"SMOE/moe/database"
	"github.com/labstack/echo/v4"
	"strings"
)

func Post(c echo.Context) error {
	qpu := database.NewQPU()
	defer qpu.Free()
	cid, err := validateNum(c.Param("cid"))
	if err != nil {
		return err
	}
	err = qpu.GetPostWithCid(cid)
	if err != nil {
		return err
	}
	err = qpu.CommentsWithCid(cid)
	if err != nil {
		return err
	}
	//fmt.Println(data.Post)
	if !strings.Contains(c.Request().Header.Get(echo.HeaderAccept), echo.MIMETextHTML) {
		return c.Render(200, "post_ajax.template", qpu)
	}
	return c.Render(200, "post.template", qpu)
}
