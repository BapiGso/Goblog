package blog

import (
	"SMOE/moe/database"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"strings"
)

// SubmitArticleComment todo 工作量证明
func SubmitArticleComment(c echo.Context) error {
	req := struct {
		Author string `xml:"author"   form:"author" validate:"required,min=1,max=200"`
		Mail   string `xml:"mail"     form:"mail" validate:"email,required,min=1,max=200"`
		Text   string `xml:"text"     form:"text" validate:"required,min=1,max=1000"`
		Url    string `xml:"url"      form:"url" validate:"omitempty,url,min=1,max=200" `
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if !strings.HasPrefix(c.Request().Referer(), c.Request().Header.Get("Origin")+"/archives/"+c.Param("cid")) {
		return echo.NewHTTPError(400, "请从评论区提交评论")
	}
	b, err := json.Marshal(&req)
	if err != nil {
		return err
	}
	m := make(map[string]any)
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	if err = database.InsertComment(m); err != nil {
		return err
	}
	return c.JSON(200, "success")
}
