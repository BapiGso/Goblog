package moe

import (
	"SMOE/moe/admin"
	"SMOE/moe/blog"
	"SMOE/moe/customw"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"text/template"
)

func (s *Smoe) LoadMiddlewareRoutes() {
	s.e.Validator = &customw.Validator{}
	s.e.Renderer = &customw.TemplateRender{
		Template: template.Must(template.ParseFS(s.ThemeFS, "*/*.template")),
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
		Level: 3,
	}))
	//http重定向https
	//e.Pre(middleware.HTTPSRedirect())
	//302跳转去除尾部斜杠
	s.e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
	//自定义404

	s.e.HTTPErrorHandler = blog.FrontErr //自定义404

	// 前台页面路由
	s.e.StaticFS("/assets", s.ThemeFS)                            // 静态文件路由，指向主题的文件系统，例如CSS，图片等静态资源
	s.e.Static("/usr/uploads", "/usr/uploads")                    //用户上传的文件
	s.e.GET("/", blog.Index)                                      // 首页路由
	s.e.GET("/page/:num", blog.Index)                             // 分页路由，显示指定页数的文章列表
	s.e.POST("/page/:num", blog.IndexAjax)                        // 分页路由，通过异步请求更新指定页数的文章列表
	s.e.GET("/archives", blog.Archive)                            // 归档页面路由，显示所有文章的归档分类
	s.e.GET("/archives/:cid", blog.Post)                          // 根据分类ID显示该分类下的文章列表
	s.e.POST("/archives/:cid/comment", blog.SubmitArticleComment) // 管理评论提交

	// 番剧页面路由
	s.e.GET("/bangumi", blog.Bangumi) // 显示番剧相关信息的页面路由

	// 后台管理
	g := s.e.Group("/admin")                                                // 后台管理的路由组
	g.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))    // 使用 session 中间件
	g.Use(customw.IsLogin)                                                  // 用户登录验证中间件
	g.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10))) //每秒10次请求
	// 后台管理页面路由
	g.GET("", admin.LoginGet)                                                     // 后台管理登录页面路由
	g.POST("", admin.LoginPost)                                                   // 后台管理登录处理路由
	g.GET("/manage-posts", admin.ManagePost, customw.SetDefaultQueryParams)       // 显示文章管理界面的路由
	g.GET("/manage-pages", admin.ManagePage)                                      // 显示页面管理界面的路由
	g.GET("/manage-comments", admin.ManageComment, customw.SetDefaultQueryParams) // 显示评论管理界面的路由
	g.GET("/manage-medias", admin.ManageMedia)                                    // 显示媒体管理界面的路由
	g.GET("/write-post", admin.WritePost)                                         // 显示添加文章页面的路由
	g.GET("/write-page", admin.WritePage)                                         // 显示添加页面页面的路由
	g.GET("/test", admin.Test)                                                    // 显示文章管理界面的路由
	g.GET("/log-access", admin.LogAccess, customw.SetDefaultQueryParams)
	g.GET("/setting", admin.Setting)

	// 文件上传路由
	g.POST("/uploadImage", admin.UploadImage) // 处理图片上传请求的路由
	g.GET("/uploadtest", admin.UploadTest)    // 文件上传测试路由，用于测试文件上传服务是否正常
}
