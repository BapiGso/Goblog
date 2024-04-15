package database

import (
	"SMOE/moe/tools"
	"bytes"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"
)

type Contents struct {
	Cid          int    `db:"cid"`
	Mid          int    `db:"mid"`   //分类用的
	Title        string `db:"title"` //用*string代替可能为null的值
	Slug         string `db:"slug"`
	Created      int64  `db:"created"`
	Text         string `db:"text"`
	Type         string `db:"type"`
	Status       string `db:"status"`
	AllowComment uint   `db:"allowComment"`
	AllowFeed    uint   `db:"allowFeed"`
	Views        uint   `db:"views"`
	Likes        uint   `db:"likes"`
	CoverList    string `db:"coverList"`
	MusicList    string `db:"musicList"`
}

// MD2HTML markdown转换为html
func (c Contents) MD2HTML() string {
	var buf bytes.Buffer
	_ = tools.GoldMark.Convert(*(*[]byte)(unsafe.Pointer(&c.Text)), &buf)
	return buf.String()
}

// MDSub 截取前95个字符串作为摘要
func (c Contents) MDSub() string {
	length := len([]rune(c.Text))
	if length <= 70 {
		return c.Text
	}
	r := string([]rune(c.Text)[:70])
	return r
}

// MDCount 计算文章字数
func (c Contents) MDCount() int {
	r := utf8.RuneCount(*(*[]byte)(unsafe.Pointer(&c.Text)))
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
