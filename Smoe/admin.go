package Smoe

import (
	"github.com/BapiGso/SMOE/smoe/query"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

type loginReq struct {
	Name     string `xml:"user"  form:"user" `
	Pwd      string `xml:"pwd" form:"pwd" `
	Illsions string `xml:"illsions"  form:"illsions" `
}

func (s *Smoe) LoginGet(c echo.Context) error {
	sess, _ := session.Get("smoesession", c)
	if sess.Values["isLogin"] == true {
		return c.Render(http.StatusOK, "admin.template", nil)
	}
	return c.Render(http.StatusOK, "login.template", nil)
}

// todo 防爆破
// todo monitor
func (s *Smoe) LoginPost(c echo.Context) error {
	req := new(loginReq)
	//调用echo.Context的Bind函数将请求参数和User对象进行绑定。
	if err := c.Bind(req); err != nil {
		return c.JSON(200, err)
	}
	sess, _ := session.Get("smoesession", c)
	//TODO 发邮件提醒和防爆破
	for _, v := range query.QueryUser(s.Db) {
		if v.Name == req.Name && v.Password == Hash(req.Pwd+v.AuthCode) {
			sess.Values["isLogin"] = true
		}
	}
	_ = sess.Save(c.Request(), c.Response())
	return c.Redirect(302, "/admin")
}
