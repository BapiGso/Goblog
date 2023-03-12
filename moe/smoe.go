package moe

import (
	"SMOE/assets"
	"SMOE/moe/mail"
	"SMOE/moe/mdparse"
	"embed"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
	"text/template"
)

type (
	Smoe struct {
		CommandLineArgs BindFlag          //命令行参数
		Db              *sqlx.DB          //数据库
		ThemeFS         *embed.FS         //主题所在文件夹
		MDParse         goldmark.Markdown //markdown->html解析器
		E               *echo.Echo        //后台框架
		Main            *mail.Email       //邮件提醒
		//异地多活
		//图片压缩webp
	}

	TemplateRender struct {
		Template *template.Template //渲染模板
	}
)

const (
	banner = `
 ______     __    __     ______     ______    
/\  ___\   /\ \-./  \   /\  __ \   /\  ___\   
\ \___  \  \ \ \-./\ \  \ \ \/\ \  \ \  __\   
 \/\_____\  \ \_\ \ \_\  \ \_____\  \ \_____\ 
  \/_____/   \/_/  \/_/   \/_____/   \/_____/ 

____________________________________O/_______
                                    O\
%s
`
)

func New() (s *Smoe) {
	s = &Smoe{}
	s.ThemeFS = &assets.Assets
	s.MDParse = mdparse.Goldmark
	s.E = echo.New()
	return s
}
