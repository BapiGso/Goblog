package moe

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "modernc.org/sqlite"
)

func (s *Smoe) InitializeDatabase() {
	var err error
	if *s.Param.DbConf != "" {
		s.Db, err = sqlx.Connect("mysql", *s.Param.DbConf)
		if err != nil {
			log.Fatalf("连接mysql数据库失败，请检查配置是否正确%v\n", err)
		}
	}
	log.Info("You are not set mysql parameter,using sqlite mode")

	s.Db, err = sqlx.Connect("sqlite", "usr/smoe.db")
	if err != nil {
		log.Fatalf("创建sqlite数据库失败，请检查读写权限%v\n", err)
	}
	log.Info("Successfully connected to the database")
	//读取sql文件创建表
	sqlTable, err := s.ThemeFS.ReadFile("smoe.sql")
	if err != nil {
		log.Fatalf("读取sql文件失败，请检查读写权限%v\n", err)
	}
	s.Db.Exec(string(sqlTable))
}
