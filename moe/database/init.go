package database

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "modernc.org/sqlite"
)

//go:embed smoe.sql
var sqlTable string

var db *sqlx.DB

func InitDB() *sqlx.DB {
	var err error
	//if *s.Param.DbConf != "" {
	//	s.Db, err = sqlx.Connect("mysql", *s.Param.DbConf)
	//	if err != nil {
	//		log.Fatalf("连接mysql数据库失败，请检查配置是否正确%v\n", err)
	//	}
	//}
	log.Info("You are not set mysql parameter,using sqlite mode")

	db, err = sqlx.Connect("sqlite", "usr/smoe.db")
	if err != nil {
		log.Fatalf("创建sqlite数据库失败，请检查读写权限%v\n", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("连接数据库失败")
	}
	log.Info("Successfully connected to the database")
	//读取sql文件创建表
	_, err = db.Exec(sqlTable)
	if err != nil {
		log.Error("创建表结构失败")
	}
	return db
}
