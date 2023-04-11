package query

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// PostWithCid 根据Cid查询单条文章或独立页面
func PostWithCid(Db *sqlx.DB, cid int) Contents {
	data := make([]Contents, 0, 1)
	_ = Db.Select(&data, `SELECT * FROM typecho_contents WHERE cid=?`, cid)
	return data[0]
}

// TestQueryPostWithCid  测试是否是指针变量
func TestQueryPostWithCid(Db *sqlx.DB, cid int) Contents {
	var data Contents
	_ = Db.Get(&data, `SELECT * FROM typecho_contents WHERE cid=?`, cid)
	return data
}

// PostArr 根据条件查询多条文章 状态 条数 页数
func PostArr(Db *sqlx.DB, status string, limit, pagenum int) []Contents {
	data := make([]Contents, 0, limit)
	_ = Db.Select(&data, `SELECT * FROM  typecho_contents
		WHERE type='post' AND status=?
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, status, limit, pagenum*limit-limit)
	return data
}

// TestAffairs  测试事务
func TestAffairs(Db *sqlx.DB, status string, limit, pagenum int) ([]Contents, []Contents) {
	var data, data2 []Contents
	tx, _ := Db.Beginx()
	go tx.Select(&data, `SELECT * FROM  typecho_contents 
		WHERE type='page' AND status=? 
		ORDER BY 'order' DESC 
		LIMIT ? OFFSET ?`, status, limit, pagenum*limit-limit)
	tx.Select(&data2, `SELECT * FROM  typecho_contents 
		WHERE type='post'
		ORDER BY 'order' `)
	return data, data2
}

// PageArr 根据条件查询多条页面
func PageArr(Db *sqlx.DB) []Contents {
	var data []Contents
	_ = Db.Select(&data, `SELECT * FROM  typecho_contents 
		WHERE type='page'
		ORDER BY 'order' `)
	return data
}

// CommentsWithCid 根据文章cid查询该文章的评论
func CommentsWithCid(Db *sqlx.DB, cid int) []Comments {
	var data []Comments
	_ = Db.Select(&data, `SELECT * FROM  typecho_comments 
		WHERE cid=?
		ORDER BY created`, cid)
	return data
}

// CommentsArr 查询评论组，后台专用
func CommentsArr(Db *sqlx.DB, status string, limit, pagenum int) []Comments {
	data := make([]Comments, 0, limit)
	_ = Db.Select(&data, `SELECT c.*,title
    	FROM typecho_comments AS c 
        INNER JOIN typecho_contents on typecho_contents.cid=c.cid
		WHERE c.status=? 
		ORDER BY c.created DESC 
		LIMIT ? OFFSET ?`, status, limit, pagenum*limit-limit)
	return data
}

// 查询文件组，后台专用
func Media(Db *sqlx.DB, limit, pagenum int) []Contents {
	data := make([]Contents, 0, limit)
	_ = Db.Select(&data, `SELECT * FROM  typecho_contents
		WHERE type='attachment'
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, limit, pagenum*limit-limit)
	return data
}

func Count(Db *sqlx.DB, Type, status string) int {
	var data int
	_ = Db.Select(&data, `SELECT count(1) FROM  typecho_contents 
		WHERE type=? AND status=?`, Type, status)
	return data
}

func UserWithName(Db *sqlx.DB, name string) (User, error) {
	var data []User
	err := Db.Select(&data, `SELECT * FROM  typecho_users WHERE name = ?`, name)
	return data[0], err
}

func InsertComment(Db *sqlx.DB, data Comments) {

}

func HaveCid(Db *sqlx.DB, cid int) bool {
	var data int
	err := Db.Get(&data, `SELECT 'allowComment' FROM typecho_contents WHERE cid = ?`, cid)
	if err != nil {
		// 如果查询过程中发生错误，可以打印错误信息并返回 false
		fmt.Printf("Error checking for CID: %v\n", err)
		return false
	}
	if data == 0 {
		return false
	}
	return true
}
