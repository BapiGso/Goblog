package main

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

type loginReq struct {
	Username string `xml:"user" json:"user" form:"user" query:"user"`
	Password string `xml:"pwd" json:"pwd" form:"pwd" query:"pwd"`
	Illsions string `xml:"illsions" json:"illsions" form:"illsions" query:"illsions"`
}

type loginSql struct {
	Username string
	Password string
	Salt     string
}

func queryLogin() *loginSql {
	data := new(loginSql)
	_ = db.QueryRow("SELECT name,password,authCode FROM typecho_users WHERE screenName='Smoe'").Scan(&data.Username, &data.Password, &data.Salt)
	return data
}

func LoginGet(c echo.Context) error {
	sess, _ := session.Get("smoesession", c)
	if sess.Values["isLogin"] == true {
		return c.Render(http.StatusOK, "admin.template", nil)
	}
	return c.Render(http.StatusOK, "login.template", nil)
}

//TODO 防爆破
func LoginPost(c echo.Context) error {
	req := new(loginReq)
	//调用echo.Context的Bind函数将请求参数和User对象进行绑定。
	if err := c.Bind(req); err != nil {
		return c.String(200, "表单提交错误")
	}
	sess, _ := session.Get("smoesession", c)
	//TODO 发邮件提醒和防爆破
	data := queryLogin()
	if data.Username == req.Username && data.Password == hash(req.Password+data.Salt) {
		sess.Values["isLogin"] = true
	} else {
		sess.Values["isLogin"] = false
	}
	sess.Save(c.Request(), c.Response())
	return c.Redirect(302, "/admin")
}
