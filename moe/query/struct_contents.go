package query

import (
	"bytes"
	"database/sql"
	"github.com/BapiGso/SMOE/moe/mdparse"
	"strings"
	"time"
	"unicode/utf8"
)

type Contents struct {
	Cid          int            `db:"cid"`
	Title        string         `db:"title"`
	Slug         string         `db:"slug"`
	Created      int64          `db:"created"`
	Modified     int64          `db:"modified"`
	Text         []byte         `db:"text"`
	Order        uint8          `db:"order"`
	AuthorId     uint8          `db:"authorId"`
	Template     sql.NullString `db:"template"`
	Type         string         `db:"type"`
	Status       string         `db:"status"`
	Password     sql.NullString `db:"password"`
	AllowComment uint8          `db:"allowComment"`
	AllowPing    uint8          `db:"allowPing"`
	AllowFeed    uint8          `db:"allowFeed"`
	CommentsNum  uint16         `db:"commentsNum"`
	Parent       uint16         `db:"parent"`
	Views        uint16         `db:"views"`
	Likes        uint32         `db:"likes"`
}

var (
	mon = map[string]string{
		"01": "一月",
		"02": "二月",
		"03": "三月",
		"04": "四月",
		"05": "五月",
		"06": "六月",
		"07": "七月",
		"08": "八月",
		"09": "九月",
		"10": "十月",
		"11": "十一月",
		"12": "十二月",
	}
)

// MD2HTML markdown转换为html
func (c Contents) MD2HTML() string {
	var buf bytes.Buffer
	_ = mdparse.Goldmark.Convert(c.Text, &buf)
	return buf.String()
}

// MDSub 截取前95字符串作为摘要
func (c Contents) MDSub() string {
	text := string(c.Text)
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
	format := (time.Unix(c.Created, 0)).Format("01 02, 2006")
	tmp := strings.Replace(format, format[:2], mon[format[:2]], 1)
	return tmp
}

func (c Contents) UnixFormat() string {
	format := (time.Unix(c.Created, 0)).Format("2006年01月02日")
	return format
}
