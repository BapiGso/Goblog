package blog

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"smoe/moe/database"
)

func Post(c echo.Context) error {
	qpu := database.NewQPU()
	defer qpu.Free()
	cid, _ := isNum(c.Param("cid"))
	//data := struct {
	//	Post     query.Contents
	//	TestPost query.Contents
	//	Comms    [][]query.Comments
	//}{
	//	query.GetPostWithCid(db, cid),
	//	query.TestQueryPostWithCid(db, cid),
	//	sortComms(query.CommentsWithCid(db, cid)),
	//}
	err := qpu.GetPostWithCid(cid)
	if err != nil {
		return err
	}
	err = qpu.CommentsWithCid(cid)
	if err != nil {
		return err
	}

	//fmt.Println(data.Post)
	return c.Render(http.StatusOK, "post.template", nil)
}
