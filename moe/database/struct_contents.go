package database

import (
	"SMOE/moe/tools"
	"bytes"
	"log/slog"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"
)

type Contents struct {
	Cid          int     `db:"cid"`
	Title        string  `db:"title"` //用*string代替可能为null的值
	Slug         string  `db:"slug"`
	Created      int64   `db:"created"`
	Modified     int64   `db:"modified"`
	Text         []byte  `db:"text"`
	Order        uint    `db:"order"`
	AuthorId     uint    `db:"authorId"`
	Template     *string `db:"template"`
	Type         string  `db:"type"`
	Status       string  `db:"status"`
	Password     *string `db:"password"`
	AllowComment uint    `db:"allowComment"`
	AllowPing    uint    `db:"allowPing"`
	AllowFeed    uint    `db:"allowFeed"`
	CommentsNum  uint    `db:"commentsNum"`
	Parent       uint    `db:"parent"`
	Views        uint    `db:"views"`
	Likes        uint    `db:"likes"`
}

// MD2HTML markdown转换为html
func (c Contents) MD2HTML() string {
	var buf bytes.Buffer
	_ = tools.GoldMark.Convert(c.Text, &buf)
	return buf.String()
}

// MDSub 截取前95个字符串作为摘要
func (c Contents) MDSub() string {
	text := *(*string)(unsafe.Pointer(&c.Text))
	length := len([]rune(text))
	if length <= 70 {
		return text
	}
	r := string([]rune(text)[:70])
	return r
}

// MDCount 计算文章字数
func (c Contents) MDCount() int {
	r := utf8.RuneCount(c.Text)
	return r
}

func (c Contents) UnixToStr() string {
	monStr := [...]string{"", "一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"}
	mon := int((time.Unix(c.Created, 0)).Month())
	format := (time.Unix(c.Created, 0)).Format("01 02, 2006")
	tmp := strings.Replace(format, format[:2], monStr[mon], 1)
	return tmp
}

func (c Contents) UnixFormat() string {
	format := (time.Unix(c.Created, 0)).Format("2006年01月02日")
	return format
}

// Bytes2String 两者指向的相同的内存，改一个另外一个也会变。
func (c Contents) Bytes2String() string {
	return *(*string)(unsafe.Pointer(&c.Text))
}

// String2Bytes https://github.com/kubernetes/apiserver/blob/706a6d89cf35950281e095bb1eeed5e3211d6272/pkg/authentication/token/cache/cached_token_authenticator.go#L263-L271
func String2Bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func (c Contents) GetMusicList() string {
	var data string
	err := DB.Get(&data, `
		SELECT str_value FROM  typecho_fields 
		WHERE cid=? and name='musicList'`, c.Cid)
	if err != nil {
		slog.Error(err.Error())
	}
	return data
}

// GetCoverList TODO 数据库无数据时随机添加一个封面
func (c Contents) GetCoverList() string {
	var data string
	err := DB.Get(&data, `
		SELECT str_value FROM  typecho_fields 
		WHERE cid=? and name='coverList'`, c.Cid)
	if err != nil {
		slog.Error(err.Error())
	}
	return data
}
