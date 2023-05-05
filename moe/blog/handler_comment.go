package blog

import (
	"database/sql"
	"fmt"
	"github.com/BapiGso/SMOE/moe/query"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"strings"
)

type CommPost struct {
	Author string `xml:"author"   form:"author"`
	Mail   string `xml:"mail"     form:"mail"`
	Text   string `xml:"text"     form:"text"`
	Url    string `xml:"url"      form:"url"`
}

func CommentPost(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	cid, err := isNum(c.Param("cid"))
	if err != nil {
		return c.JSON(400, err)
	}
	if !strings.HasPrefix(c.Request().Referer(), c.Request().Header.Get("Origin")+"/archives/"+c.Param("cid")) {
		return c.String(400, "请从评论区提交评论")
	}
	//a := template.HTMLEscapeString("http://127.0.0.1/?name=<script>alert('durban，xss')</script>")
	//绑定查询参数
	req := new(CommPost)
	if err := c.Bind(req); err != nil {
		return c.JSON(403, err)
	}
	fmt.Println(req, c.QueryParam("parent"))
	//如果数据库有该cid所在的文章
	if query.HaveCid(db, cid) {
		_ = query.Comments{
			Coid:    0,
			Cid:     0,
			OwnerId: 0,
			Parent:  0,
			Created: 0,
			Author:  "",
			Mail:    "",
			Ip:      "",
			Agent:   "",
			Text:    "",
			Type:    "",
			Status:  "",
			//Title:    "",
			AuthorId: 0,
			Url:      sql.NullString{},
		}
		return c.JSON(200, "sus")
		//query.InsertComment(db, req)
	}

	return c.JSON(200, "success")
}
