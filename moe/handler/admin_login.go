package handler

import (
	"SMOE/moe/database"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginGet(c echo.Context) error {
	sess, _ := session.Get("smoeSession", c)
	if sess.Values["isLogin"] == true {
		return c.Render(http.StatusOK, "admin.template", nil)
	}
	return c.Render(http.StatusOK, "login.template", nil)
}

func LoginPost(c echo.Context) error {
	qpu := database.NewQPU()
	defer qpu.Free()
	req := struct {
		Name     string `xml:"user"  form:"user" `
		Pwd      string `xml:"pwd" form:"pwd" `
		Illsions string `xml:"illsions"  form:"illsions" `
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	sess, err := session.Get("smoeSession", c)
	if err != nil {
		return err
	}
	err = qpu.UserWithName(req.Name)
	if err != nil {
		return err
	}
	//计算提交表单的密码与盐 scrypt和数据库中密码是否一致
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(qpu.UserInfo.Password), []byte(req.Pwd)); err == nil {
		sess.Values["isLogin"] = true
	} else {
		return err
	}

	//TODO 发邮件提醒和防爆破
	_ = sess.Save(c.Request(), c.Response())
	return c.Redirect(302, "/admin")
}
