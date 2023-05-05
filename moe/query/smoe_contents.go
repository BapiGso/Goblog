package query

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// 查询文件组，后台专用
func Media(Db *sqlx.DB, limit, pageNum int) []Contents {
	data := make([]Contents, 0, limit)
	_ = Db.Select(&data, `SELECT * FROM  typecho_contents
		WHERE type='attachment'
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, limit, pageNum*limit-limit)
	return data
}

// PostWithCid 根据Cid查询单条文章或独立页面
func PostWithCid(Db *sqlx.DB, cid int) Contents {
	data := Contents{}
	err := Db.Get(&data, `SELECT * FROM typecho_contents WHERE cid=? AND type='post'`, cid)
	fmt.Println(err)
	return data
}

// PageWithCid 根据Cid查询单条文章或独立页面
func PageWithCid(Db *sqlx.DB, cid int) Contents {
	data := Contents{}
	err := Db.Get(&data, `SELECT * FROM typecho_contents WHERE cid=? AND type='page'`, cid)
	fmt.Println(err)
	return data
}

// TestQueryPostWithCid  测试是否是指针变量
func TestQueryPostWithCid(Db *sqlx.DB, cid int) Contents {
	var data Contents
	_ = Db.Get(&data, `SELECT * FROM typecho_contents WHERE cid=?`, cid)
	return data
}

// PostArr 根据条件查询多条文章 状态 条数 页数
func PostArr(Db *sqlx.DB, status string, limit, pageNum int) []Contents {
	data := make([]Contents, 0, limit)
	_ = Db.Select(&data, `SELECT * FROM  typecho_contents
		WHERE type='post' AND status=?
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, status, limit, pageNum*limit-limit)
	return data
}

// PageArr 根据条件查询多条页面
func PageArr(Db *sqlx.DB) []Contents {
	var data []Contents
	err := Db.Select(&data, `SELECT * FROM  typecho_contents 
		WHERE type='page'
		ORDER BY 'order' `)
	fmt.Println(err)
	return data
}
