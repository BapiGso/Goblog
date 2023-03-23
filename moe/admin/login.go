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
	sess, _ := session.Get("smoesession", c)
	if sess.Values["isLogin"] == true {
		return c.Render(http.StatusOK, "admin.template", nil)
	}
	return c.Render(http.StatusOK, "login.template", nil)
}

// todo 防爆破
// todo monitor
func LoginPost(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	req := new(loginReq)
	//调用echo.Context的Bind函数将请求参数和User对象进行绑定。
	if err := c.Bind(req); err != nil {
		return c.JSON(200, err)
	}
	sess, _ := session.Get("smoesession", c)
	//TODO 发邮件提醒和防爆破
	for _, v := range query.QueryUser(db) {
		if v.Name == req.Name && v.Password == Hash(req.Pwd+v.AuthCode) {
			sess.Values["isLogin"] = true
		}
	}
	_ = sess.Save(c.Request(), c.Response())
	return c.Redirect(302, "/admin")
}

// Hash 计算字符串sha1
func Hash(input string) string {
	h := sha1.New() // md5加密类似md5.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(input))
	bs := h.Sum(nil)
	h.Reset()
	passwdhash := hex.EncodeToString(bs)
	return passwdhash
}
