package moe

import (
	"SMOE/moe/handler"
	"SMOE/moe/mymiddleware"
	"fmt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"strings"
	"text/template"
	"time"
)

func (s *Smoe) LoadMiddlewareRoutes() {
	s.e.Validator = &mymiddleware.Validator{}
	s.e.Renderer = &mymiddleware.TemplateRender{
		Template: template.Must(
			template.ParseFS(
				s.themeFS,
				"blog/*.template",
				"blog/js/*.js",
				"blog/css/*.css",
				"new-admin/*.template",
			),
		),
	}

	//Secure防XSS，HSTS防中间人攻击 todo 防盗链
	s.e.Pre(middleware.SecureWithConfig(middleware.SecureConfig{
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
			go slog.Info(fmt.Sprint(v.Error),
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

	s.e.Use(mymiddleware.BrotliWithConfig(mymiddleware.BrotliConfig{
		Level: 0,
	}))

	//http重定向https
	//s.e.Pre(middleware.HTTPSRedirect())

	//301跳转去除尾部斜杠
	s.e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	//使用session
	s.e.Use(echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			//判断当前路径是否是受限制路径（除了登录页面以外的后台路径）
			restricted := strings.HasPrefix(c.Path(), "/admin") && c.Path() != "/admin"
			if _, err := c.Cookie("smoe_token"); err != nil && !restricted {
				return true
			}
			slog.Warn("someone not have token visit restricted!")
			return false
		},
		ErrorHandler: func(c echo.Context, err error) error {
			//因为我只在正确登录后发正确token,要么就是没token
			//所以只有错误token会触发该函数 todo 触发错误后ip限制
			c.SetCookie(&http.Cookie{Name: "smoe_token", Expires: time.Now(), MaxAge: -1, HttpOnly: true})
			return echo.ErrUnauthorized
		},
		//todo 成功后给qpu权限
		SigningKey:  handler.SigningKey,
		TokenLookup: "cookie:smoe_token",
	}))

	//自定义404
	s.e.HTTPErrorHandler = handler.FrontErr //自定义404
	front := s.e.Group("")
	back := s.e.Group("/admin")

	// 前台页面路由
	front.GET("/", handler.Index)                                      // 首页路由
	front.GET("/page/:num", handler.Index)                             // 分页路由，显示指定页数的文章列表
	front.GET("/archives/:cid", handler.Post)                          // 根据分类ID显示该分类下的文章列表
	front.POST("/archives/:cid/comment", handler.SubmitArticleComment) // 管理评论提交
	front.GET("/:page", handler.Page)                                  //独立页面，注册在特殊独立页面前
	front.GET("/archives", handler.Archives)                           // 归档页面路由，显示所有文章的归档分类
	front.GET("/bangumi", handler.Bangumi)                             // 显示番剧相关信息的页面路由
	front.Static("/usr/uploads", "/usr/uploads")                       //用户上传的文件，最后注册
	front.StaticFS("/assets", s.themeFS)                               // 静态文件路由,最后注册

	// 后台管理
	// 后台管理的路由组
	back.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(3))) //每秒限制3次请求
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
