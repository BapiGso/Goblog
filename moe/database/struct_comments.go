package database

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"time"
)

type Comments struct {
	Coid     uint32         `db:"coid"     form:"coid"`
	Cid      uint32         `db:"cid"      form:"cid"`
	OwnerId  uint32         `db:"ownerId"  form:"ownerId"`
	Parent   uint32         `db:"parent"   form:"parent"`
	Created  int64          `db:"created"  form:"created"`
	Author   string         `db:"author"   form:"author"`
	Mail     string         `db:"mail"     form:"mail"`
	Ip       string         `db:"ip"       form:"ip"`
	Agent    string         `db:"agent"    form:"agent"`
	Text     string         `db:"text"     form:"text"`
	Type     string         `db:"type"     form:"type"`
	Status   string         `db:"status"   form:"status"`
	AuthorId uint8          `db:"authorId" form:"authorId"`
	Url      sql.NullString `db:"url"      form:"url"`
}

func (c Comments) UnixFormat() string {
	format := (time.Unix(c.Created, 0)).Format("2006年01月02日")
	return format
}

func (c Comments) MD5Mail() string {
	data := md5.Sum([]byte(c.Mail))
	md5str := fmt.Sprintf("%x", data)
	return md5str
}

func (c Comments) SubText() string {
	// 将字符串转换为[]rune，以便正确处理Unicode字符
	runes := []rune(c.Text)
	runesLength := len(runes)

	if runesLength <= 20 {
		return c.Text
	}

	more := fmt.Sprintf(`...<a class="tooltip" data-tooltip="%v">查看更多</a>`, c.Text)
	return string(runes[:20]) + more
}
