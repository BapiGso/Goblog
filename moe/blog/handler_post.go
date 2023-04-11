package blog

import (
	"github.com/BapiGso/SMOE/moe/query"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Post(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	cid, _ := isNum(c.Param("cid"))
	data := struct {
		Post     query.Contents
		TestPost query.Contents
		Comms    [][]query.Comments
	}{
		query.PostWithCid(db, cid),
		query.TestQueryPostWithCid(db, cid),
		sortComms(query.CommentsWithCid(db, cid)),
	}
	//fmt.Println(data.Post)
	return c.Render(http.StatusOK, "post.template", data)
}
