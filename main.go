package main

import (
	"database/sql"
	"flag"
	"github.com/gorilla/sessions"
	_ "github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"log"
	"main/smoe"
	_ "main/smoe"
	_ "modernc.org/sqlite"
	"net/http"
	_ "net/http/pprof"
	"os"
	"text/template"
)

var db *sql.DB

var s = Smoe.New()

func init() {
	checkDB()

	//test()
}

func checkDB() {
	//不存在就创建数据库
	var err error
	db, err = sql.Open("sqlite", "usr/smoe.db")
	if err != nil {
		log.Fatalf("创建数据库失败，请检查读写权限%v\n", err)
	}
	//读取sql文件创建表
	sqlTable, err := os.ReadFile("usr/smoe.sql")
	if err != nil {
		log.Fatalf("读取sql文件失败，请检查读写权限%v\n", err)
	}
	db.Exec(string(sqlTable))
}

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

// todo auto tls
func main() {
	//go http.ListenAndServe(":8080", nil)
	//s := souin_echo.New(souin_echo.DevDefaultConfiguration)
	//c := freecache.NewCache(1024 * 1024 * 0)
	bind := flag.String("http", ":8081", "bind address")
	flag.Parse()
	e := echo.New()
	e.Renderer = &Smoe.TEmplateRender{
		TemplateRender: template.Must(template.ParseFS(s.ThemeFS, "*/*.template")),
	}

	//e.Logger.SetLevel(log.DEBUG)
	//Secure防XSS，HSTS防中间人攻击
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		HSTSMaxAge:            31536000,
		HSTSPreloadEnabled:    true,
		HSTSExcludeSubdomains: true,
	}))

	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//echoV5更新时换成broitil编码
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 3,
	}))
	//http重定向https
	//e.Pre(middleware.HTTPSRedirect())
	//302跳转去除尾部斜杠
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
	//自定义404
	e.HTTPErrorHandler = func(err error, c echo.Context) { c.Render(http.StatusNotFound, "404.template", err) } //自定义404
	//e.HTTPErrorHandler = errdel
	e.StaticFS("/", s.ThemeFS)
	e.GET("/", Index)
	e.GET("/page/:num", Index)
	e.POST("/page/:num", IndexAjax)
	e.GET("/archives", Archive)
	e.GET("/archives/:cid", Post)
	e.GET("/bangumi", Bangumi)
	g := e.Group("/admin")
	g.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	g.Use(IsLogin)
	g.GET("", LoginGet)
	g.POST("", LoginPost)
	g.GET("/manage-posts", ManagePost)
	g.GET("/manage-pages", ManagePage)
	g.GET("/manage-comments", ManageComment)
	g.GET("/manage-medias", ManageMedia)
	g.GET("/write-post", WritePost)
	g.GET("/write-page", WritePage)
	g.POST("/upload", Upload)
	g.GET("/uploadtest", UploadTest)
	e.Start(*bind)
}

func errdel(err error, c echo.Context) {

}
