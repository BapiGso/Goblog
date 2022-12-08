package Smoe

import (
	"crypto/sha1"
	"embed"
	"encoding/hex"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
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
	}

	TEmplateRender struct {
		TemplateRender *template.Template //渲染模板
	}
)

var (
	MDParse = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Linkify,
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

func (t *TEmplateRender) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.TemplateRender.ExecuteTemplate(w, name, data)
}

// QueryPostWithCid 根据Cid查询单条文章
func (s *Smoe) QueryPostWithCid(cid int) Contents {
	var data Contents
	_ = s.Db.Get(&data, `SELECT * FROM typecho_contents WHERE cid=?`, cid)
	return data
}

// QueryPostArr 根据条件查询多条文章 状态 条数 页数
func (s *Smoe) QueryPostArr(status string, limit, pagenum uint64) []Contents {
	var data []Contents
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
func (s *Smoe) QueryCommentsWithCid(cid int) []Comments {
	var data []Comments
	_ = s.Db.Select(&data, `SELECT * FROM  typecho_comments 
		WHERE cid=?`, cid)
	return data
}

func (s *Smoe) QueryMedia(limit, pagenum int) []Contents {
	var data []Contents
	_ = s.Db.Select(&data, `SELECT * FROM  typecho_contents 
		WHERE type='attachment'
		ORDER BY rowid DESC 
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
