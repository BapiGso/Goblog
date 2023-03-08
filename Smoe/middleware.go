package Smoe

import (
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"text/template"
)

func (s *Smoe) Middleware() {
	s.E.Renderer = &TemplateRender{
		Template: template.Must(template.ParseFS(s.ThemeFS, "*/*.template")),
	}

	//e.Logger.SetLevel(log.DEBUG)
	//Secure防XSS，HSTS防中间人攻击
	s.E.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		HSTSMaxAge:            31536000,
		HSTSPreloadEnabled:    true,
		HSTSExcludeSubdomains: true,
	}))

	//e.Use(middleware.Logger())
	s.E.Use(middleware.Recover())

	//echoV5更新时换成broitil编码
	s.E.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 3,
	}))
	//http重定向https
	//e.Pre(middleware.HTTPSRedirect())
	//302跳转去除尾部斜杠
	s.E.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
}
