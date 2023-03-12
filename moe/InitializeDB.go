package moe

import (
	"github.com/jmoiron/sqlx"
	"log"
)

func (s *Smoe) InitializeDatabase() {
	var err error
	if s.CommandLineArgs.DbConf != "" {
		s.Db, err = sqlx.Connect("mysql", s.CommandLineArgs.DbConf)
		if err != nil {
			log.Fatalf("连接数据库失败，请检查读写权限%v\n", err)
		}
	}

	s.Db, err = sqlx.Connect("sqlite", "usr/smoe.db")
	if err != nil {
		log.Fatalf("创建数据库失败，请检查读写权限%v\n", err)
	}

	//读取sql文件创建表
	sqlTable, err := s.ThemeFS.ReadFile("smoe.sql")
	if err != nil {
		log.Fatalf("读取sql文件失败，请检查读写权限%v\n", err)
	}
	s.Db.Exec(string(sqlTable))
}
