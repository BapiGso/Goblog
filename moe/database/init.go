package database

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	"log/slog"
	_ "modernc.org/sqlite"
)

//go:embed smoe.sql
var sqlTable string

var DB = func() *sqlx.DB {
	//if *s.Param.DbConf != "" {
	//	s.Db, err = sqlx.Connect("mysql", *s.Param.DbConf)
	//	if err != nil {
	//		log.Fatalf("连接mysql数据库失败，请检查配置是否正确%v\n", err)
	//	}
	//}
	db := sqlx.MustConnect("sqlite", "usr/smoe.db")

	//读取sql文件创建表
	_, err := db.Exec(sqlTable)
	if err != nil {
		slog.Error("创建表结构失败")
	}
	return db
}()
