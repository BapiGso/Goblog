package handler

import (
	"SMOE/moe/database"
	"github.com/labstack/echo/v4"
)

func Write(c echo.Context) error {
	qpu := new(database.QPU)
	req := &struct {
		Cid       string `param:"cid" `
		Slug      string `form:"slug" `
		Title     string `form:"title" `
		Text      string `form:"text"`
		CoverList string `form:"coverList"`
		MusicList string `form:"musicList"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	reqMap, err := struct2map(*req)
	if err != nil {
		return err
	}
	if reqMap["Slug"] == "" {
		reqMap["Slug"] = req.Cid
		reqMap["Type"] = "post"
	} else {
		reqMap["Type"] = "page"
	}
	switch c.Request().Method {
	case "GET": //渲染攥写文章页面
		//如果cid为new，则是写新文章
		if req.Cid == "new" {
			return c.Render(200, "write.template", nil)
		}
		cid, err := validateNum(c.Param("cid"))
		if err != nil {
			return err
		}
		if err := database.DB.Select(&qpu.Contents, `
			SELECT * FROM smoe_contents 
        	WHERE cid=?`, cid); err != nil {
			return err
		}
		return c.Render(200, "write.template", qpu)
	case "POST": //新建文章的API
		if err := database.InsertContent(reqMap); err != nil {
			return err
		}
		return c.JSON(201, nil)
	case "PUT": //更新文章的API
		if _, err := database.DB.NamedExec(`UPDATE smoe_contents
		SET title=:title,slug=:slug,text=:text
		WHERE cid=:cid`, req); err != nil {
			return err
		}
		return c.JSON(200, nil)
	case "DELETE": //todo 删除文章的API

	}

	return nil
}
