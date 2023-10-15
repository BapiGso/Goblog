package handler

import (
	"SMOE/moe/database"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"time"
)

//todo complete

type TimelineData struct {
	Cid         uint64
	CreatedUnix int64
	Day         int
	Mon         int
	Year        int
	Title       string
}

type TimelineTmpdata struct {
	Cid   uint64
	Day   int
	Title string
}

type h1 struct {
	Mon   int
	MData []TimelineTmpdata //Mdata里有这个月的文章
}

type h2 struct {
	Year  int
	YData []h1 //Ydata里有12个月
}

type h3 struct {
	Title string
	Data  []h2
}

func QueryTime(Db *sqlx.DB) h3 {
	//data := make([]TimelineData, 95)
	data := TimelineData{}
	//var timeline = map[int]map[int][]TimelineTmpdata{}
	rows, _ := Db.Query(`SELECT r.cid,c.title,c.created
		FROM typecho_relationships AS r 
		INNER JOIN typecho_contents AS c ON c.cid=r.cid 
		WHERE c.status='publish'
		ORDER BY r.cid DESC `)
	//for i := 0; rows.Next(); i++ {
	//	_ = rows.Scan(&data[i].Cid, &data[i].Title, &data[i].CreatedUnix)
	//	data[i].Year = (time.Unix(data[i].CreatedUnix, 0)).Year()
	//	data[i].Mon = int((time.Unix(data[i].CreatedUnix, 0)).Month())
	//	if data[i].Year != 0 {
	//		if timeline[data[i].Year] == nil {
	//			timeline[data[i].Year] = make(map[int][]TimelineTmpdata, 5)
	//		}
	//		//fmt.Println(v.Year)
	//		timeline[data[i].Year][data[i].Mon] = append(timeline[data[i].Year][data[i].Mon], TimelineTmpdata{data[i].Cid, data[i].Title})
	//	}
	//}
	//rows.Close()
	var lasty, lastm int
	aa1, aa2, aa3 := h1{}, h2{}, h3{} //一个月的数据,一年的数据，总数据
	for i := 0; rows.Next(); i++ {
		_ = rows.Scan(&data.Cid, &data.Title, &data.CreatedUnix)
		data.Year = (time.Unix(data.CreatedUnix, 0)).Year()
		data.Mon = int((time.Unix(data.CreatedUnix, 0)).Month())
		data.Day = (time.Unix(data.CreatedUnix, 0)).Day()
		if i == 0 {
			lasty, lastm = data.Year, data.Mon
		}
		if data.Mon == lastm { //如果月份和上次添加的一样
			aa1.MData = append(aa1.MData, TimelineTmpdata{data.Cid, data.Day, data.Title}) //添加文章到该月份
		}
		if data.Mon != lastm && data.Year == lasty { //如果月份和上次添加的不一样
			aa2.YData = append(aa2.YData, h1{lastm, aa1.MData})                            //将该月份数据添加到该年
			aa1.MData = nil                                                                //该月份数据清空
			aa1.MData = append(aa1.MData, TimelineTmpdata{data.Cid, data.Day, data.Title}) //添加文章到该月份
		}
		if data.Year != lasty {
			aa2.YData = append(aa2.YData, h1{lastm, aa1.MData})
			aa1.MData = nil
			aa1.MData = append(aa1.MData, TimelineTmpdata{data.Cid, data.Day, data.Title})
			aa3.Data = append(aa3.Data, h2{lasty, aa2.YData}) //将该年数据添加到总数据
			aa2.YData = nil                                   //该年数据清空
		}
		lastm, lasty = data.Mon, data.Year
		//fmt.Println(i)
	}
	//退出循环后把最后一次数据加上
	aa2.YData = append(aa2.YData, h1{lastm, aa1.MData})
	aa3.Data = append(aa3.Data, h2{lasty, aa2.YData}) //将该年数据添加到总数据
	aa3.Title = "时间线"
	rows.Close()
	//fmt.Println(aa3.Data)
	//判断年份是否存在，存在就判断月份，没有就添加，判断月份，没有添加，有就append文章
	//var timeline = map[int]map[int][]TimelineTmpdata{}
	//for _, v := range data {
	//	if v.Year != 0 {
	//		if timeline[v.Year] == nil {
	//			timeline[v.Year] = make(map[int][]TimelineTmpdata, 5)
	//		}
	//		//fmt.Println(v.Year)
	//		timeline[v.Year][v.Mon] = append(timeline[v.Year][v.Mon], TimelineTmpdata{v.Cid, v.Title})
	//	}
	//}
	//fmt.Println(timeline)
	return aa3
}

func Archive(c echo.Context) error {
	qpu := database.NewQPU()
	defer qpu.Free()
	return echo.NewHTTPError(400, "")
	return c.Render(200, "page-timeline.template", qpu)
}
