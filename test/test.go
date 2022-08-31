package main

import (
	"github.com/labstack/echo/v4"
)

func init() {
	test2()
}

func queryTest(test [10]TestData) {
	test[0].aa = 123
	//fmt.Println(test[0].aa)
}

type TestData struct {
	aa uint64
	//PageMax  uint64
}

type bb struct {
	cc [10]TestData
}

func main() {
	e := echo.New()
	e.GET("/", test)
	//slice := bb{}
	//slice := make([]int, 0, 5)
	//fmt.Println(slice)
	e.Start(":8082")
}

func test(c echo.Context) error {
	ha := &bb{}
	queryTest(ha.cc)
	//fmt.Println(ha)
	return c.Render(200, "test.template", ha)
}

func test2() {
	hah := map[int]map[int][]TestData{}
	//hah[int] = make(map[int][]TestData{})
	hah[1][1] = append(hah[1][1], TestData{123})
	hah[1][1] = append(hah[1][1], TestData{56})
	println(hah[1][1][1].aa)
}
