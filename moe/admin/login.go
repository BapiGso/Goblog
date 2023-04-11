package admin

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/BapiGso/SMOE/moe/query"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

type loginReq struct {
	Name     string `xml:"user"  form:"user" `
	Pwd      string `xml:"pwd" form:"pwd" `
	Illsions string `xml:"illsions"  form:"illsions" `
}

func LoginGet(c echo.Context) error {
	sess, _ := session.Get("smoeSession", c)
	if sess.Values["isLogin"] == true {
		return c.Render(http.StatusOK, "admin.template", nil)
	}
	return c.Render(http.StatusOK, "login.template", nil)
}

func LoginPost(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	req := new(loginReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, err)
	}
	sess, err := session.Get("smoeSession", c)
	if err != nil {
		return c.JSON(400, err)
	}
	user, err := query.UserWithName(db, req.Name)
	if err != nil {
		return c.JSON(400, err)
	}
	if user.Password == hash(req.Pwd+user.AuthCode) {
		sess.Values["isLogin"] = true
	}
	//TODO 发邮件提醒和防爆破
	_ = sess.Save(c.Request(), c.Response())
	return c.Redirect(302, "/admin")
}

// hash 计算字符串sha1
func hash(input string) string {
	h := sha1.New() // md5加密类似md5.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(input))
	bs := h.Sum(nil)
	h.Reset()
	passwdhash := hex.EncodeToString(bs)
	return passwdhash
}
