package query

import "github.com/jmoiron/sqlx"

// CommentsWithCid 根据文章cid查询该文章的评论
func CommentsWithCid(Db *sqlx.DB, cid int) []Comments {
	var data []Comments
	_ = Db.Select(&data, `SELECT * FROM  typecho_comments 
		WHERE cid=?
		ORDER BY created`, cid)
	return data
}
