package Smoe

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/BapiGso/SMOE/smoe/mdparse"
	_ "modernc.org/sqlite"
	"strings"
	"time"
	"unicode/utf8"
)

// 内存对齐 https://geektutu.com/post/hpg-struct-alignment.html
type (
	Comments struct {
		Coid     uint32         `db:"coid"`
		Cid      uint32         `db:"cid"`
		OwnerId  uint32         `db:"ownerId"`
		Parent   uint32         `db:"parent"`
		Created  int64          `db:"created"`
		Author   string         `db:"author"`
		Mail     string         `db:"mail"`
		Ip       string         `db:"ip"`
		Agent    string         `db:"agent"`
		Text     string         `db:"text"`
		Type     string         `db:"type"`
		Status   string         `db:"status"`
		Title    string         `db:"title"`
		AuthorId uint8          `db:"authorId"`
		Url      sql.NullString `db:"url"`
	}

	Contents struct {
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

	Fields struct {
		Cid        int     `db:"cid"`
		Name       string  `db:"name"`
		Type       string  `db:"type"`
		StrValue   string  `db:"str_value"`
		IntValue   int     `db:"int_value"`
		FloatValue float64 `db:"float_value"`
	}

	User struct {
		Uid        string `db:"uid"`
		Name       string `db:"name"`
		Password   string `db:"password"`
		Mail       string `db:"mail"`
		Url        string `db:"url"`
		ScreenName string `db:"screenName"`
		Created    int64  `db:"created"`
		Activated  int64  `db:"activated"`
		Logged     int64  `db:"logged"`
		Group      string `db:"group"`
		AuthCode   string `db:"authCode"`
	}
)

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
func (c *Contents) MD2HTML() string {
	var buf bytes.Buffer
	_ = mdparse.Goldmark.Convert(c.Text, &buf)
	return buf.String()
}

// MDSub 截取前95字符串作为摘要
func (c *Contents) MDSub() string {
	r := string([]rune(string(c.Text))[:70])
	return r
}

// MDCount 计算文章字数
func (c *Contents) MDCount() int {
	r := utf8.RuneCount(c.Text)
	return r
}

func (c *Contents) UnixToStr() string {
	format := (time.Unix(c.Created, 0)).Format("01 02, 2006")
	tmp := strings.Replace(format, format[:2], mon[format[:2]], 1)
	return tmp
}

func (c *Contents) UnixFormat() string {
	format := (time.Unix(c.Created, 0)).Format("2006年01月02日")
	return format
}

func (c *Comments) UnixFormat() string {
	format := (time.Unix(c.Created, 0)).Format("2006年01月02日")
	return format
}

func (c *Comments) MD5Mail() string {
	data := md5.Sum([]byte(c.Mail))
	md5str := fmt.Sprintf("%x", data)
	return md5str
}
