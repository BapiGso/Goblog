package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type PostManageData struct {
	CommNum     uint32
	Cid         uint32
	CreatedUnix int64
	CreatedStr  string
	Title       string
}

//又是一个大坑，数据库里有NULL的用sql.NullString
type CommManageData struct {
	Coid        uint16
	Cid         uint16
	CreatedUnix int64
	CreatedStr  string
	Author      string
	Mail        string
	MailMD5     string
	Url         sql.NullString
	Ip          string
	Text        string
	Title       string
}

type MediaManageData struct {
	Cid         uint16
	CreatedUnix int64
	CreatedStr  string
	Text        string
	Title       string
	Parent      uint16
}

//类名首字母一定要大写，又被坑一次
type ManageParam struct {
	Status string `query:"status" form:"status" json:"status"`
	Page   string `query:"page" form:"page" json:"page"`
}

func md5v(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func queryManageMedia(data *[20]MediaManageData, pageNum uint64, limit uint64) {
	rows, _ := db.Query(`SELECT cid,title,created,text,parent 
		FROM typecho_contents 
		WHERE type='attachment' 
		ORDER BY rowid DESC LIMIT ? OFFSET ?`, limit, pageNum*limit-limit)
	for i := 0; rows.Next(); i++ {
		_ = rows.Scan(&data[i].Cid, &data[i].Title, &data[i].CreatedUnix, &data[i].Text, &data[i].Parent)
		data[i].CreatedStr = unix2time(data[i].CreatedUnix)
		//fmt.Println(data.Cid)
	}
	rows.Close()
}

func queryManageComment(data *[20]CommManageData, status string, pageNum uint64, limit uint64) {
	rows, _ := db.Query(`SELECT c.coid,c.cid,c.created,c.author,c.mail,c.url,c.ip,c.text,title 
		FROM typecho_comments AS c 
		INNER JOIN typecho_contents on typecho_contents.cid=c.cid 
		WHERE c.status=? 
		ORDER BY c.rowid DESC LIMIT ? OFFSET ?`, status, limit, pageNum*limit-limit)
	for i := 0; rows.Next(); i++ {
		_ = rows.Scan(&data[i].Coid, &data[i].Cid, &data[i].CreatedUnix, &data[i].Author, &data[i].Mail, &data[i].Url, &data[i].Ip, &data[i].Text, &data[i].Title)
		data[i].CreatedStr = unix2time(data[i].CreatedUnix)
		data[i].MailMD5 = md5v(data[i].Mail)
	}
	rows.Close()
}

func queryManagePost(data *[20]PostManageData, postStatus string, pageNum uint64, limit uint64) {
	rows, _ := db.Query(`SELECT r.cid,c.title,c.commentsNum,c.created 
		FROM typecho_relationships AS r 
    	INNER JOIN typecho_contents AS c on c.cid=r.cid 
		WHERE c.status=? 
		ORDER BY r.cid DESC LIMIT ? OFFSET ?`, postStatus, limit, pageNum*limit-limit)
	for i := 0; rows.Next(); i++ {
		_ = rows.Scan(&data[i].Cid, &data[i].Title, &data[i].CommNum, &data[i].CreatedUnix)
		data[i].CreatedStr = (time.Unix(data[i].CreatedUnix, 0)).Format("2006-01-02 15:04")
	}
	rows.Close()
}

func ManagePost(c echo.Context) error {
	req := new(ManageParam)
	req.Status = "publish"
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	pageNum, _ := isNum(req.Page)
	postSlice := [20]PostManageData{}
	queryManagePost(&postSlice, req.Status, pageNum, 20)
	//fmt.Println(postSlice)
	//还要查询页数
	return c.Render(200, "manage-posts.template", &postSlice)
}

func ManagePage(c echo.Context) error {
	pageSlice := [10]PageData{}
	queryPage(&pageSlice)
	//fmt.Println(postSlice[0].Title)
	return c.Render(200, "manage-pages.template", &pageSlice)
}

func ManageComment(c echo.Context) error {
	req := new(ManageParam)
	req.Status = "approved"
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	pageNum, _ := isNum(req.Page)
	commSlice := [20]CommManageData{}
	queryManageComment(&commSlice, req.Status, pageNum, 20)
	//fmt.Println(commSlice[0].Title)
	return c.Render(200, "manage-comments.template", commSlice)
}

func ManageMedia(c echo.Context) error {
	req := new(ManageParam)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "参数Param错误")
	}
	pageNum, _ := isNum(req.Page)
	mediaSlice := [20]MediaManageData{}
	queryManageMedia(&mediaSlice, pageNum, 20)
	return c.Render(200, "manage-medias.template", mediaSlice)
}
