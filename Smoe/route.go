package Smoe

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
)

func (s *Smoe) LoadRoutes() {
	// 前台页面路由
	s.E.StaticFS("/", s.ThemeFS)            // 静态文件路由，指向主题的文件系统，例如CSS，图片等静态资源
	s.E.GET("/", s.BlogIndex)               // 首页路由
	s.E.GET("/page/:num", s.BlogIndexAjax)  // 分页路由，显示指定页数的文章列表
	s.E.POST("/page/:num", s.BlogIndexAjax) // 分页路由，通过异步请求更新指定页数的文章列表
	// 归档分类页面路由
	s.E.GET("/archives", s.Archive)   // 归档页面路由，显示所有文章的归档分类
	s.E.GET("/archives/:cid", s.Post) // 根据分类ID显示该分类下的文章列表

	// 番剧页面路由
	s.E.GET("/bangumi", s.Bangumi) // 显示番剧相关信息的页面路由

	// 后台管理路由
	g := s.E.Group("/admin")                                             // 后台管理的路由组
	g.Use(session.Middleware(sessions.NewCookieStore([]byte("secret")))) // 使用 session 中间件
	g.Use(IsLogin)                                                       // 用户登录验证中间件

	// 后台管理页面路由
	g.GET("", s.LoginGet)                      // 后台管理登录页面路由
	g.POST("", s.LoginPost)                    // 后台管理登录处理路由，接收登录表单的提交
	g.GET("/manage-posts", s.ManagePost)       // 显示文章管理界面的路由
	g.GET("/manage-pages", s.ManagePage)       // 显示页面管理界面的路由
	g.GET("/manage-comments", s.ManageComment) // 显示评论管理界面的路由
	g.GET("/manage-medias", s.ManageMedia)     // 显示媒体管理界面的路由
	g.GET("/write-post", s.WritePost)          // 显示添加文章页面的路由
	g.GET("/write-page", s.WritePage)          // 显示添加页面页面的路由

	// 文件上传路由
	g.POST("/upload", s.Upload)        // 处理文件上传请求的路由
	g.GET("/uploadtest", s.UploadTest) // 文件上传测试路由，用于测试文件上传服务是否正常

}
