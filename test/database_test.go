package test

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
	"testing"
)

func TestSQLiteConnection(t *testing.T) {
	// 连接 SQLite 数据库
	db, err := sqlx.Connect("sqlite", "../usr/smoe.db")
	if err != nil {
		t.Fatalf("连接 SQLite 数据库失败：%v", err)
	}
	defer db.Close()

	// 测试是否连接成功
	err = db.Ping()
	if err != nil {
		t.Fatalf("连接 SQLite 数据库失败：%v", err)
	}

	t.Logf("连接 SQLite 数据库成功")
}
