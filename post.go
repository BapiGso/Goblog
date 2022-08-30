package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PostData struct {
	Cid         uint16
	Likes       uint16
	Views       uint32
	CreatedUnix uint32
	TextLen     uint32
	Title       string
	Cover       string
	Text        string
	Text2HTML   string
}

func queryArchive(data *PostData, cid uint64) {
	_ = db.QueryRow(`SELECT typecho_contents.cid,title,created,text,length(text),views,likes,str_value 
		FROM typecho_contents  
		inner join typecho_fields on typecho_contents.cid=typecho_fields.cid 
		WHERE typecho_contents.cid=? and typecho_contents.status='publish' 
		and typecho_fields.name='bgMusicList' `, cid).Scan(&data.Cid, &data.Title, &data.CreatedUnix, &data.Text, &data.TextLen, &data.Views, &data.Likes, &data.Cover)
}

//func queryComment(cid uint64, status string, limit uint64) []CommData {
//	commSlice := make([]CommData, 0, limit)
//	data := CommData{}
//	rows, _ := db.Query("SELECT coid,cid,created,author,authorId,mail,url,ip,text,parent "+
//		"FROM typecho_comments  "+
//		"WHERE cid=(?) and status=? ", cid, status)
//	defer rows.Close()
//	for rows.Next() {
//		_ = rows.Scan(&data.Coid, &data.Cid, &data.Created, &data.Author, &data.AuthorId, &data.Mail, &data.Url, &data.Ip, &data.Text, &data.Parent)
//		commSlice = append(commSlice, data)
//	}
//	return commSlice
//}

//TODO 查不到的抛出404

func Post(c echo.Context) error {
	cid, _ := isNum(c.Param("cid"))
	data := PostData{}
	queryArchive(&data, cid)
	fmt.Println(data.Cid)
	return c.Render(http.StatusOK, "post.template", data)
}
