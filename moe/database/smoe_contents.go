package database

import (
	"errors"
)

// Media 查询文件组，后台专用
func (q *QPU) Media(limit, pageNum int) error {
	err := db.Select(&q.PostArr, `
		SELECT * FROM  typecho_contents
		WHERE type='attachment'
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, limit, pageNum*limit-limit)
	return err
}

// GetPostWithCid 根据Cid查询文章 todo 权限
func (q *QPU) GetPostWithCid(cid int) error {
	err := db.Select(&q.PostArr, `SELECT * FROM typecho_contents WHERE cid=? AND type='post'`, cid)
	if len(q.PostArr) == 0 {
		return errors.New("没有查询到结果")
	}
	return err
}

// GetPageWithCid 根据Cid查询独立页面
func (q *QPU) GetPageWithCid(cid int) error {
	err := db.Select(&q.PageArr, `SELECT * FROM typecho_contents WHERE cid=? AND type='page'`, cid)
	return err
}

// GetPosts  根据条件查询多条文章 状态 条数 页数
func (q *QPU) GetPosts(status string, limit, pageNum int) error {
	err := db.Select(&q.PostArr, `
		SELECT * FROM  typecho_contents
		WHERE type='post' AND status=?
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, status, limit+1, pageNum*limit-limit)
	//多查一个判断是否有下一页
	if len(q.PostArr) == 6 {
		q.HaveNext = pageNum + 1
		q.PostArr = q.PostArr[:len(q.PostArr)-1]
	}
	return err
}

// GetPage  根据条件查询单独页面
func (q *QPU) GetPage(p string) error {
	err := db.Select(&q.PageArr, `
		SELECT * FROM  typecho_contents 
		WHERE type='page' AND slug = ?`, p)
	if len(q.PageArr) == 0 {
		return errors.New("not found")
	}
	return err
}

// GetPages  根据条件查询多条页面
func (q *QPU) GetPages() error {
	err := db.Select(&q.PageArr, `
		SELECT * FROM  typecho_contents 
		WHERE type='page'
		ORDER BY "order" `)
	return err
}

// CheckNext 查询是否有下一页，如果有则ture
func (q *QPU) CheckNext() error {
	err := db.Get(q.HaveNext, `
		SELECT cid FROM  typecho_contents 
		WHERE cid>?`, q.PostArr[len(q.PostArr)-1].Cid)
	return err
}

func UpdateView(cid string) error {
	_, err := db.Exec(`UPDATE typecho_contents SET views = views + 1 WHERE cid = ?`, cid)
	if err != nil {
		return err
	}
	return nil
}
