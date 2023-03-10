package query

import "database/sql"

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
