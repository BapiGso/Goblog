package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

//类名首字母一定要大写，又被坑一次
type WriteParam struct {
	Cid uint16 `query:"cid" form:"cid" json:"cid"`
}

type WriteData struct {
	Cid         uint32
	Order       uint32
	CreatedUnix int64
	Modified    int64
	Slug        string
	CreatedStr  string
	Title       string
	Text        string
	Music       string
	Cover       string
}

func WriteQuery(data *WriteData, cid uint16) {
	//fmt.Println(cid)
	_ = db.QueryRow(`SELECT cid,slug,title,created,modified,substr(text,16),'order' FROM typecho_contents WHERE cid=?`, cid).Scan(&data.Cid, &data.Slug, &data.Title, &data.CreatedUnix, &data.Modified, &data.Text, &data.Order)
	_ = db.QueryRow(`SELECT str_value FROM typecho_fields WHERE name='coverList'   AND cid=?`, cid).Scan(&data.Cover)
	_ = db.QueryRow(`SELECT str_value FROM typecho_fields WHERE name='bgMusicList' AND cid=?`, cid).Scan(&data.Music)
	data.CreatedStr = (time.Unix(data.CreatedUnix, 0)).Format("2006-01-02 15:04")
}

func WritePost(c echo.Context) error {
	req := new(WriteParam)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	if req.Cid == 0 {
		return c.Render(200, "write-post.template", nil)
	}
	data := new(WriteData)
	WriteQuery(data, req.Cid)
	//fmt.Println(data)
	return c.Render(200, "write-post.template", data)

}

func WritePage(c echo.Context) error {
	req := new(WriteParam)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	if req.Cid == 0 {
		return c.Render(200, "write-page.template", nil)
	}
	data := new(WriteData)
	WriteQuery(data, req.Cid)
	//fmt.Println(data)
	return c.Render(200, "write-page.template", data)
}
