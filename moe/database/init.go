package database

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	"log/slog"
	_ "modernc.org/sqlite"
)

//go:embed smoe.sql
var sqlTable string

var db *sqlx.DB

func InitDB() {
	var err error
	//if *s.Param.DbConf != "" {
	//	s.Db, err = sqlx.Connect("mysql", *s.Param.DbConf)
	//	if err != nil {
	//		log.Fatalf("连接mysql数据库失败，请检查配置是否正确%v\n", err)
	//	}
	//}
	slog.Info("You are not set mysql parameter,using sqlite mode")

	if db, err = sqlx.Connect("sqlite", "usr/smoe.db"); err != nil {
		slog.Error("连接sqlite数据库失败，请检查读写权限", err)
	}
	if err := db.Ping(); err != nil {
		slog.Error("连接数据库失败")
	}

	slog.Info("Successfully connected to the database")
	//读取sql文件创建表
	_, err = db.Exec(sqlTable)
	if err != nil {
		slog.Error("创建表结构失败")
	}
}
