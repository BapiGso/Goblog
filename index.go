package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type PageData struct {
	Cid   uint64
	Title string
	Slug  string
}

type IndexPostData struct {
	Views       uint16
	Likes       uint16
	Cid         uint32 //路径
	TextLen     uint32 //字数
	CreatedUnix int64  //时间戳
	Title       string
	TextSub     string //截取
	CreatedStr  string //十二月 23,2022
	CoverMusic  string //封面或者音乐
}

// IndexData template模板渲染的类要首字母大写，坑了我半天
type IndexData struct {
	IndexPost [5]IndexPostData
	PageItem  [10]PageData
	PageNext  uint64
	//PageMax  uint64
}

func queryMaxPage() uint64 {
	var num uint64
	_ = db.QueryRow(`SELECT count(cid) from typecho_contents where type='post' AND status='publish'`).Scan(&num)
	//println(num)
	return num
}

//查询页面的路由路径，返回一个slice切片
func queryPage(data *[10]PageData) {
	rows, _ := db.Query(`SELECT cid,title,slug FROM typecho_contents WHERE type='page' ORDER BY "order"`)
	for i := 0; rows.Next(); i++ {
		_ = rows.Scan(&data[i].Cid, &data[i].Title, &data[i].Slug)
		//fmt.Println(data[i].Title)
	}
	rows.Close()
}

func queryPost(data *[5]IndexPostData, postStatus string, pageNum uint64, limit uint64) {
	//查询一篇文章的数据在加上封面，总共查5篇数据
	rows, _ := db.Query(`SELECT Slug,title,created,substr(text,0,70),length(text),views,likes,str_value 
		FROM typecho_contents  
		inner join typecho_fields on typecho_contents.cid=typecho_fields.cid 
		WHERE typecho_contents.type='post' and typecho_contents.status=? and typecho_fields.name='coverList' 
		ORDER BY typecho_contents.cid DESC LIMIT ? OFFSET ?`, postStatus, limit, pageNum*limit-limit)
	for i := 0; rows.Next(); i++ {
		_ = rows.Scan(&data[i].Cid, &data[i].Title, &data[i].CreatedUnix, &data[i].TextSub, &data[i].TextLen, &data[i].Views, &data[i].Likes, &data[i].CoverMusic)
		data[i].CreatedStr = unix2time(data[i].CreatedUnix)
		//fmt.Println(data[i].Title)
	}
	rows.Close()
}

func Index(c echo.Context) error {
	//判断页数查数据库
	pageNum, _ := isNum(c.Param("num"))
	//postMaxNum := queryMaxPage()
	//if pageNum > postMaxNum/5+1 {
	//	return c.Render(http.StatusNotFound, "404.template", nil)
	//}
	//缓存是否存在
	//_, ok := indexCache[pageNum]
	//if ok {
	//	return c.Render(http.StatusOK, "index.template", indexCache[pageNum])
	//}
	//fmt.Printf("index%lu\n", unsafe.Sizeof(IndexData{}))
	indexData := new(IndexData)
	queryPage(&indexData.PageItem)
	queryPost(&indexData.IndexPost, "publish", pageNum, 5)
	indexData.PageNext = pageNum + 1
	//indexSlice.PageMax = postMaxNum/5 + 2
	//indexCache[pageNum] = indexData
	//go 访客增加views
	//reCacheIndex <- true
	return c.Render(http.StatusOK, "index.template", indexData)
}
