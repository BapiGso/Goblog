package Smoe

import (
	"database/sql"
	"log"
)

func init() {
	checkDB()

	//test()
}

func checkDB() {

}

func (s *Smoe) InitializeDatabase() {
	//不存在就创建数据库
	var err error
	db, err := sql.Open("sqlite", "usr/smoe.db")
	if err != nil {
		log.Fatalf("创建数据库失败，请检查读写权限%v\n", err)
	}
	//读取sql文件创建表
	sqlTable, err := s.ThemeFS.ReadFile("smoe.sql")
	if err != nil {
		log.Fatalf("读取sql文件失败，请检查读写权限%v\n", err)
	}
	db.Exec(string(sqlTable))
}
