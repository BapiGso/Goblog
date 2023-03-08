package main

import (
	_ "github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"io"
	"main/smoe"
	_ "main/smoe"
	_ "modernc.org/sqlite"
	"net/http"
	_ "net/http/pprof"
	"text/template"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func IsLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//登录页面不用这个中间件
		if c.Path() == "/admin" {
			return next(c)
		}
		//后台页面没有cookie的全部跳去登录
		sess, _ := session.Get("smoesession", c)
		if sess.Values["isLogin"] != true {
			return c.Redirect(http.StatusFound, "/admin")
		}
		return next(c)
	}
}

func main() {
	s := Smoe.New()
	s.BindFlag()
	s.Middleware()
	s.LoadRoutes()
	s.Listen()
}
