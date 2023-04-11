package query

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"time"
)

type Comments struct {
	Coid     uint32         `db:"coid"`
	Cid      uint32         `db:"cid"`
	OwnerId  uint32         `db:"ownerId"`
	Parent   uint32         `db:"parent"`
	Created  int64          `db:"created"`
	Author   string         `db:"author"`
	Mail     string         `db:"mail" xml:"mail"  form:"mail" `
	Ip       string         `db:"ip"`
	Agent    string         `db:"agent"`
	Text     string         `db:"text" xml:"author" form:"author"`
	Type     string         `db:"type"`
	Status   string         `db:"status"`
	Title    string         `db:"title"`
	AuthorId uint8          `db:"authorId"`
	Url      sql.NullString `db:"url" xml:"url"  form:"url"`
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
