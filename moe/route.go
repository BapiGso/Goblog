package moe

import (
	"SMOE/moe/customw"
	"SMOE/moe/handler"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"strings"
	"text/template"
)

func (s *Smoe) LoadMiddlewareRoutes() {
	s.e.Validator = &customw.Validator{}
	s.e.Renderer = &customw.TemplateRender{
		Template: template.Must(template.ParseFS(s.themeFS, "*/*.template")),
	}

	//Secure防XSS，HSTS防中间人攻击
	s.e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		HSTSMaxAge:            31536000,
		HSTSPreloadEnabled:    true,
		HSTSExcludeSubdomains: true,
	}))
	//s.e.Logger.SetLevel(log.INFO)
	s.e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:   true,
		LogURI:      true,
		LogStatus:   true,
		LogRemoteIP: true,
		LogError:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			slog.Info(fmt.Sprint(v.Error),
				slog.String("Method", v.Method),
				slog.String("Url", v.URI),
				slog.Int("Status", v.Status),
				slog.String("IP", v.RemoteIP))
			return nil
		},
	}))
	//s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	//echoV5更新时换成broitil编码
	s.e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			//浏览器支持br的时候跳过使用gzip压缩
			return strings.Contains(c.Request().Header.Get(echo.HeaderAcceptEncoding), "br")
		},
		Level: 3,
	}))
	s.e.Use(customw.BrotliWithConfig(customw.BrotliConfig{
		Level: 0,
	}))

	//http重定向https
	//e.Pre(middleware.HTTPSRedirect())
	//302跳转去除尾部斜杠
	s.e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	//自定义404
	s.e.HTTPErrorHandler = handler.FrontErr //自定义404

	front := s.e.Group("")
	back := s.e.Group("/admin")
	// 前台页面路由

	front.GET("/", handler.Index)          // 首页路由
	front.GET("/page/:num", handler.Index) // 分页路由，显示指定页数的文章列表
	//todo 不跟据请求方法，根据req header是否有x request
	front.POST("/page/:num", handler.IndexAjax)                        // 分页路由，通过异步请求更新指定页数的文章列表
	front.GET("/archives/:cid", handler.Post)                          // 根据分类ID显示该分类下的文章列表
	front.POST("/archives/:cid/comment", handler.SubmitArticleComment) // 管理评论提交
	front.GET("/:page", handler.Page)                                  //独立页面，注册在特殊独立页面前
	front.GET("/archives", handler.Archive)                            // 归档页面路由，显示所有文章的归档分类
	front.GET("/bangumi", handler.Bangumi)                             // 显示番剧相关信息的页面路由
	front.Static("/usr/uploads", "/usr/uploads")                       //用户上传的文件，最后注册
	front.StaticFS("/assets", s.themeFS)                               // 静态文件路由,最后注册

	// 后台管理
	// 后台管理的路由组
	back.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))    // 使用 session 中间件
	back.Use(customw.CheckAuthMiddleware)                                      // 用户登录验证中间件
	back.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10))) //每秒10次请求
	// 后台管理页面路由
	back.GET("", handler.LoginGet)                      // 后台管理登录页面路由
	back.POST("", handler.LoginPost)                    // 后台管理登录处理路由
	back.GET("/manage-posts", handler.ManagePost)       // 显示文章管理界面的路由
	back.GET("/manage-pages", handler.ManagePage)       // 显示页面管理界面的路由
	back.GET("/manage-comments", handler.ManageComment) // 显示评论管理界面的路由
	back.GET("/manage-medias", handler.ManageMedia)     // 显示媒体管理界面的路由
	back.GET("/write-post", handler.WritePost)          // 显示添加文章页面的路由
	back.GET("/write-page", handler.WritePage)          // 显示添加页面页面的路由
	back.GET("/test", handler.Test)                     // 显示文章管理界面的路由
	back.GET("/log-access", handler.LogAccess)
	back.GET("/setting", handler.Setting)

	// 文件上传路由
	back.POST("/uploadImage", handler.UploadImage) // 处理图片上传请求的路由
	back.GET("/uploadtest", handler.UploadTest)    // 文件上传测试路由，用于测试文件上传服务是否正常
}
