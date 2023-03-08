package Smoe

import (
	"embed"
	mermaid "github.com/abhinav/goldmark-mermaid"
	latex "github.com/aziis98/goldmark-latex"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"io"
	"log"
	"main/assets"
	"text/template"
)

type (
	Smoe struct {
		CommandLineArgs BindFlag          //命令行参数
		Db              *sqlx.DB          //数据库
		ThemeFS         *embed.FS         //主题所在文件夹
		MDParse         goldmark.Markdown //markdown->html解析器
		E               *echo.Echo        //后台框架
		//邮件提醒
		//异地多活
		//图片压缩webp
	}

	TemplateRender struct {
		Template *template.Template //渲染模板
	}
)

var (
	MDParse = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Linkify,
			//mathjax.MathJax,
			&mermaid.Extender{},
			//latex.NewLatex(
			//	latex.WithSourceInlineDelim(`\(`, `\)`),
			//	latex.WithSourceBlockDelim(`\[`, `\]`),
			//	latex.WithOutputInlineDelim(`\(`, `\)`),
			//	latex.WithOutputBlockDelim(`\[`, `\]`),
			//),
			highlighting.NewHighlighting(
				highlighting.WithStyle("github")),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
		),
	)
)

func New() (s *Smoe) {
	var err error
	s = &Smoe{}
	s.Db, err = sqlx.Connect("sqlite", "smoe/smoe.db")
	s.ThemeFS = &assets.Assets
	if err != nil {
		log.Fatalf("创建数据库失败，请检查读写权限%v\n", err)
	}
	sqltable, _ := s.ThemeFS.ReadFile("smoe.sql")
	_, _ = s.Db.Exec(string(sqltable))

	s.MDParse = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			mathjax.MathJax,
			&mermaid.Extender{},
			latex.NewLatex(),
			highlighting.NewHighlighting(
				highlighting.WithStyle("github")),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
		),
	)
	s.E = echo.New()
	return s
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.Template.ExecuteTemplate(w, name, data)
}
