package main

import "github.com/labstack/echo/v4"

func queryTest(data []TestData) {
	db.QueryRow(`SELECT count(*) FROM typecho_contents WHERE type='page' ORDER BY "order"`).Scan(&data[0].aa)
	//fmt.Println(test[0].aa)
}

func queryTest2(data []TestData) {
	db.QueryRow(`SELECT count(*) FROM typecho_contents WHERE type='page' ORDER BY "order"`).Scan(&data[1].aa)
	//fmt.Println(test[0].aa)
}

func queryTest3() uint64 {
	var data uint64
	db.QueryRow(`SELECT count(*) FROM typecho_contents WHERE type='page' ORDER BY "order"`).Scan(&data)
	return data
	//fmt.Println(test[0].aa)
}

type TestData struct {
	aa uint64
	//PageMax  uint64
}

type bb struct {
	cc []TestData
}

func test(c echo.Context) error {
	ha := bb{
		make([]TestData, 5),
	}
	queryTest(ha.cc)
	queryTest2(ha.cc)
	ha.cc[0].aa = queryTest3()
	//fmt.Println(ha)
	habixia := ha
	return c.Render(200, "404.template", habixia)
}
