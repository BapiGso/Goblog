package moe

import (
	"SMOE/moe/handler"
	"SMOE/moe/mymiddleware"
	"context"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
	"path/filepath"
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
		).Funcs(template.FuncMap{}),
	}
	//Secure防XSS，HSTS防中间人攻击 todo 防盗链
	s.e.Pre(middleware.SecureWithConfig(middleware.SecureConfig{
		HSTSMaxAge:            31536000,
		HSTSPreloadEnabled:    true,
		HSTSExcludeSubdomains: true,
	}))
	//cors防盗链
	//s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"http://localhost:8080"}, // 允许的源地址
	//	AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	//}))
	// 中间件：禁用跨站请求伪造（CSRF）
	//s.e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	//	TokenLookup: "header:X-CSRF-Token", // 从请求头中获取CSRF令牌
	//}))

	//s.e.Logger.SetLevel(log.INFO)
	s.e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:   true,
		LogURI:      true,
		LogStatus:   true,
		LogRemoteIP: true,
		LogError:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				slog.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("Method", v.Method),
					slog.String("Url", v.URI),
					slog.String("IP", v.RemoteIP),
					slog.Int("Status", v.Status),
				)
			} else {
				slog.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("Method", v.Method),
					slog.String("Url", v.URI),
					slog.Int("Status", v.Status),
					slog.String("IP", v.RemoteIP),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	//s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	//echoV5更新时换成broitil编码
	s.e.Use(mymiddleware.BrotliWithConfig(mymiddleware.BrotliConfig{
		Level: 0,
	}))
	s.e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			//如果调用了brotli，那么就跳过Gzip
			return c.Response().Header().Get(echo.HeaderContentEncoding) == "br"
		},
		Level: 3,
	}))

	//http重定向https
	//s.e.Pre(middleware.HTTPSRedirect())

	//301跳转去除尾部斜杠
	s.e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	//使用jwt
	s.e.Use(echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			restricted := strings.HasPrefix(c.Path(), "/admin/") //判断当前路径是否是受限制路径（除了登录页面以外的后台路径）
			_, err := c.Cookie("smoe_token")
			return err != nil && !restricted //如果读不到cookie且不是受限制路径就跳过
		},
		ErrorHandler: func(c echo.Context, err error) error {
			//todo 触发错误后ip限制
			c.SetCookie(&http.Cookie{Name: "smoe_token", Expires: time.Now(), MaxAge: -1, HttpOnly: true})
			return echo.ErrTeapot
		},
		SuccessHandler: func(c echo.Context) {

		},
		SigningKey:  mymiddleware.JWTKey,
		TokenLookup: "cookie:smoe_token",
	}))
	s.e.Group("/assets", middleware.StaticWithConfig(middleware.StaticConfig{
		//skipper跳过一些不想让用户和爬虫看到的文件
		Skipper: func(c echo.Context) bool {
			ext := filepath.Ext(c.Request().URL.Path)
			return !(ext == ".css" || ext == ".js" || ext == ".ico" || ext == ".svg" || ext == ".webp")
		},
		Filesystem: http.FS(s.themeFS),
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
	front.Static("/usr/uploads", "usr/uploads")                        //用户上传的文件，最后注册

	// 后台管理
	// 后台管理的路由组
	back.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(3))) //每秒限制3次请求
	// 后台管理页面路由
	back.GET("", handler.LoginGet)   // 后台管理登录页面路由
	back.POST("", handler.LoginPost) // 后台管理登录处理路由
	back.Any("/write/:cid", handler.Write)
	back.Any("/manage/:type", handler.Manage)
	back.GET("/log-access", handler.LogAccess)
	back.GET("/setting", handler.Setting)

	// 文件上传路由
	back.POST("/uploadImage", handler.UploadImage) // 处理图片上传请求的路由
	back.GET("/uploadtest", handler.UploadTest)    // 文件上传测试路由，用于测试文件上传服务是否正常
}
