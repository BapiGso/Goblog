package moe

import (
	"SMOE/moe/mail"
	"embed"
	"github.com/jmoiron/sqlx"
	"smoe/assets"
	_ "smoe/moe/database"
	"text/template"
)

type (
	Smoe struct {
		Param   *BindFlag //命令行参数
		Db      *sqlx.DB  //数据库
		ThemeFS *embed.FS //主题所在文件夹
		//mdParse *goldmark.Markdown //markdown->html解析器
		e    *echo.Echo  //后台框架
		Mail *mail.Email //邮件提醒
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
	s.e = echo.New()
	return s
}
