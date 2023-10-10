package customw

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//登录页面不用这个中间件
		if c.Path() == "/admin" {
			return next(c)
		}
		//后台页面没有cookie的全部跳去登录
		sess, err := session.Get("smoeSession", c)
		if err != nil {
			return err
		}
		if sess.Values["isLogin"] != true {
			return c.Redirect(http.StatusFound, "/admin")
		}
		return next(c)
	}
}

func test(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		return next(c)
	}
}
