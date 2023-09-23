package customw

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func IsLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//登录页面不用这个中间件
		if c.Path() == "/admin" {
			return next(c)
		}
		//后台页面没有cookie的全部跳去登录
		sess, _ := session.Get("smoeSession", c)
		if sess.Values["isLogin"] != true {
			return c.Redirect(http.StatusFound, "/admin")
		}
		return next(c)
	}
}

func SetDefaultQueryParams(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		queryParams := map[string]string{
			"commstatus": "approved",
			"status":     "publish",
			"page":       "1",
		}
		for key, value := range queryParams {
			if c.QueryParam(key) == "" {
				c.QueryParams().Set(key, value)
			}
		}
		return next(c)
	}
}
