package blog

import (
	"github.com/BapiGso/SMOE/moe/query"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func CommentPost(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	cid, err := isNum(c.Param("cid"))
	if err != nil {
		return c.JSON(400, err)
	}
	req := new(query.Comments)
	if err := c.Bind(req); err != nil {
		return c.JSON(403, err)
	}
	//如果数据库有该cid所在的文章
	if query.HaveCid(db, cid) {
		return c.JSON(200, "sus")
		//检查该评论的referrer是否一致，还有其他安全措施，然后写入数据库
		//query.InsertComment(db, req)
	}
	return c.JSON(200, "success")
}
