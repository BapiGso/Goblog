package Smoe

import (
	"crypto/sha1"
	"embed"
	"encoding/hex"
	"fmt"
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
	"strconv"
	"text/template"
	"time"
)

type (
	Smoe struct {
		Db      *sqlx.DB          //数据库
		ThemeFS *embed.FS         //主题所在文件夹
		MDParse goldmark.Markdown //markdown->html解析器
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
	return
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.Template.ExecuteTemplate(w, name, data)
}

// QueryWithCid 根据Cid查询单条文章或独立页面
func (s *Smoe) QueryWithCid(cid uint64) []Contents {
	data := make([]Contents, 0, 1)
	_ = s.Db.Select(&data, `SELECT * FROM typecho_contents WHERE cid=?`, cid)
	return data
}

// TestQueryPostWithCid  测试是否是指针变量
func (s *Smoe) TestQueryPostWithCid(cid uint64) Contents {
	var data Contents
	_ = s.Db.Get(&data, `SELECT * FROM typecho_contents WHERE cid=?`, cid)
	return data
}

// QueryPostArr 根据条件查询多条文章 状态 条数 页数
func (s *Smoe) QueryPostArr(status string, limit, pagenum uint64) []Contents {
	data := make([]Contents, 0, limit)
	_ = s.Db.Select(&data, `SELECT * FROM  typecho_contents 
		WHERE type='post' AND status=? 
		ORDER BY ROWID DESC 
		LIMIT ? OFFSET ?`, status, limit, pagenum*limit-limit)
	return data
}

// Testshiwu 测试事务
func (s *Smoe) Testshiwu(status string, limit, pagenum uint64) ([]Contents, []Contents) {
	var data, data2 []Contents
	tx, _ := s.Db.Beginx()
	go tx.Select(&data, `SELECT * FROM  typecho_contents 
		WHERE type='page' AND status=? 
		ORDER BY 'order' DESC 
		LIMIT ? OFFSET ?`, status, limit, pagenum*limit-limit)
	tx.Select(&data2, `SELECT * FROM  typecho_contents 
		WHERE type='post'
		ORDER BY 'order' `)
	return data, data2
}

// QueryPageArr 根据条件查询多条页面
func (s *Smoe) QueryPageArr() []Contents {
	var data []Contents
	_ = s.Db.Select(&data, `SELECT * FROM  typecho_contents 
		WHERE type='page'
		ORDER BY 'order' `)
	return data
}

// QueryCommentsWithCid 根据文章cid查询该文章的评论
func (s *Smoe) QueryCommentsWithCid(cid uint64) []Comments {
	var data []Comments
	_ = s.Db.Select(&data, `SELECT * FROM  typecho_comments 
		WHERE cid=?`, cid)
	return data
}

// QueryCommentsArr 查询评论组，后台专用
func (s *Smoe) QueryCommentsArr(status string, limit, pagenum uint64) []Comments {
	data := make([]Comments, 0, limit)
	_ = s.Db.Select(&data, `SELECT c.*,title
    	FROM typecho_comments AS c 
        INNER JOIN typecho_contents on typecho_contents.cid=c.cid
		WHERE c.status=? 
		ORDER BY c.created DESC 
		LIMIT ? OFFSET ?`, status, limit, pagenum*limit-limit)
	return data
}

// 查询文件组，后台专用
func (s *Smoe) QueryMedia(limit, pagenum uint64) []Contents {
	data := make([]Contents, 0, limit)
	_ = s.Db.Select(&data, `SELECT * FROM  typecho_contents
		WHERE type='attachment'
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, limit, pagenum*limit-limit)
	return data
}

func (s *Smoe) QueryCount(Type, status string) uint64 {
	var data uint64
	_ = s.Db.Select(&data, `SELECT count(1) FROM  typecho_contents 
		WHERE type=? AND status=?`, Type, status)
	return data
}

func (s *Smoe) QueryUser() []User {
	var data []User
	_ = s.Db.Select(&data, `SELECT * FROM  typecho_users`)
	return data
}

// Hash 计算字符串sha1
func Hash(input string) string {
	h := sha1.New() // md5加密类似md5.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(input))
	bs := h.Sum(nil)
	h.Reset()
	passwdhash := hex.EncodeToString(bs)
	return passwdhash
}

// IsNum 首页返回1，不是数字返回err调用404，其他为对应页数
func IsNum(numstr string) (uint64, error) {
	if numstr == "" {
		return 1, nil
	}
	num, err := strconv.ParseUint(numstr, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func TimeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("time cost = %v\n", tc)
	}
}
