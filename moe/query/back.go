package query

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

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

// 给定当前页数和一页的数量，查询剩余页数
func NextPageNum(Db *sqlx.DB, status string, limit, pageNum int) int {
	var data int
	_ = Db.Get(&data, `SELECT count(1) FROM  typecho_contents 
		WHERE type='post' AND status=?`, status)
	if data > limit*pageNum {
		return pageNum + 1
	}
	return 0
}

func InsertComment(Db *sqlx.DB, data Comments) {

}

func HaveCid(Db *sqlx.DB, cid int) bool {
	var data int
	err := Db.Get(&data, `SELECT allowComment FROM typecho_contents WHERE cid = ?`, cid)
	if err != nil {
		// 如果查询过程中发生错误，可以打印错误信息并返回 false
		fmt.Print(err)
		return false
	}
	return data != 0
}
