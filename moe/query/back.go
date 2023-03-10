package query

import "github.com/jmoiron/sqlx"

// QueryWithCid 根据Cid查询单条文章或独立页面
func QueryWithCid(Db *sqlx.DB, cid uint64) []Contents {
	data := make([]Contents, 0, 1)
	_ = Db.Select(&data, `SELECT * FROM typecho_contents WHERE cid=?`, cid)
	return data
}

// TestQueryPostWithCid  测试是否是指针变量
func TestQueryPostWithCid(Db *sqlx.DB, cid uint64) Contents {
	var data Contents
	_ = Db.Get(&data, `SELECT * FROM typecho_contents WHERE cid=?`, cid)
	return data
}

// QueryPostArr 根据条件查询多条文章 状态 条数 页数
func QueryPostArr(Db *sqlx.DB, status string, limit, pagenum uint64) []Contents {
	data := make([]Contents, 0, limit)
	_ = Db.Select(&data, `SELECT * FROM  typecho_contents 
		WHERE type='post' AND status=? 
		ORDER BY ROWID DESC 
		LIMIT ? OFFSET ?`, status, limit, pagenum*limit-limit)
	return data
}

// TestAffairs  测试事务
func TestAffairs(Db *sqlx.DB, status string, limit, pagenum uint64) ([]Contents, []Contents) {
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

// QueryPageArr 根据条件查询多条页面
func QueryPageArr(Db *sqlx.DB) []Contents {
	var data []Contents
	_ = Db.Select(&data, `SELECT * FROM  typecho_contents 
		WHERE type='page'
		ORDER BY 'order' `)
	return data
}

// QueryCommentsWithCid 根据文章cid查询该文章的评论
func QueryCommentsWithCid(Db *sqlx.DB, cid uint64) []Comments {
	var data []Comments
	_ = Db.Select(&data, `SELECT * FROM  typecho_comments 
		WHERE cid=?`, cid)
	return data
}

// QueryCommentsArr 查询评论组，后台专用
func QueryCommentsArr(Db *sqlx.DB, status string, limit, pagenum uint64) []Comments {
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
func QueryMedia(Db *sqlx.DB, limit, pagenum uint64) []Contents {
	data := make([]Contents, 0, limit)
	_ = Db.Select(&data, `SELECT * FROM  typecho_contents
		WHERE type='attachment'
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, limit, pagenum*limit-limit)
	return data
}

func QueryCount(Db *sqlx.DB, Type, status string) uint64 {
	var data uint64
	_ = Db.Select(&data, `SELECT count(1) FROM  typecho_contents 
		WHERE type=? AND status=?`, Type, status)
	return data
}

func QueryUser(Db *sqlx.DB) []User {
	var data []User
	_ = Db.Select(&data, `SELECT * FROM  typecho_users`)
	return data
}
