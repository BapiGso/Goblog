package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type PostData struct {
	Cid         uint16
	Likes       uint16
	Views       uint32
	TextLen     uint32
	CreatedUnix int64
	CreatedStr  string
	Title       string
	Music       string
	Text        []byte
	Text2HTML   string
	Comm        [][]CommentData
}

type CommentData struct {
	Cid         uint16
	AuthorID    uint16
	Parent      uint32
	Coid        uint32
	CreatedUnix int64
	Author      string
	CreatedStr  string
	Mail        string
	MailMD5     string
	Url         sql.NullString
	Text        string
}

func queryArchive(data *PostData, cid uint64) {
	_ = db.QueryRow(`SELECT typecho_contents.cid,title,created,text,length(text),views,likes,str_value 
		FROM typecho_contents  
		inner join typecho_fields on typecho_contents.cid=typecho_fields.cid 
		WHERE typecho_contents.cid=? and typecho_contents.status='publish' 
		and typecho_fields.name='bgMusicList' `, cid).Scan(&data.Cid, &data.Title, &data.CreatedUnix, &data.Text, &data.TextLen, &data.Views, &data.Likes, &data.Music)
	data.Text2HTML = md2html(data.Text)
	data.CreatedStr = unix2time(data.CreatedUnix)
}

//TODO 倒序
func queryComment(data *PostData, cid uint64) {
	rows, _ := db.Query(`SELECT cid,coid,created,author,authorId,mail,url,text,parent
		FROM typecho_comments
		WHERE cid=? AND status='approved'`, cid)
	for rows.Next() {
		tmpdata := CommentData{}
		_ = rows.Scan(&tmpdata.Cid, &tmpdata.Coid, &tmpdata.CreatedUnix, &tmpdata.Author, &tmpdata.AuthorID, &tmpdata.Mail, &tmpdata.Url, &tmpdata.Text, &tmpdata.Parent)
		tmpdata.MailMD5 = md5v(tmpdata.Mail)
		tmpdata.CreatedStr = (time.Unix(tmpdata.CreatedUnix, 0)).Format("2006-01-02 15:04")
		//fmt.Println(tmpdata.Parent)
		if tmpdata.Parent == 0 { //如果是新的一组评论
			data.Comm = append(data.Comm, []CommentData{})                             //加一组评论
			data.Comm[len(data.Comm)-1] = append(data.Comm[len(data.Comm)-1], tmpdata) //把这条评论加到改组中
			//fmt.Println("有新评论")
		} else { //如果是回复评论
			for k1, v1 := range data.Comm {
				for _, v2 := range v1 {
					if tmpdata.Parent == v2.Coid {
						data.Comm[k1] = append(data.Comm[k1], tmpdata) //遍历找父评论，加到父评论的组中
					}
				}
			}
		}

	}
	//fmt.Println(data[3])
	rows.Close()
}

//TODO 查不到文章的抛出404

func Post(c echo.Context) error {
	cid, _ := isNum(c.Param("cid"))
	data := PostData{}
	queryArchive(&data, cid)
	queryComment(&data, cid)
	if data.Cid == 0 {
		return echo.ErrNotFound
	}
	//fmt.Println(data.Comm[3])
	return c.Render(http.StatusOK, "post.template", data)
}
