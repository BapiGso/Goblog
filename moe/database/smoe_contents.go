package database

import (
	"errors"
)

// Media 查询文件组，后台专用
func (s *QPU) Media(limit, pageNum int) error {
	err := db.Select(&s.PostArr, `
		SELECT * FROM  typecho_contents
		WHERE type='attachment'
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, limit, pageNum*limit-limit)
	return err
}

// GetPostWithCid 根据Cid查询文章
func (s *QPU) GetPostWithCid(cid int) error {
	err := db.Select(&s.PostArr, `SELECT * FROM typecho_contents WHERE cid=? AND type='post'`, cid)
	if len(s.PostArr) == 0 {
		return errors.New("没有查询到结果")
	}
	return err
}

// GetPageWithCid 根据Cid查询独立页面
func (s *QPU) GetPageWithCid(cid int) error {
	err := db.Select(&s.PageArr, `SELECT * FROM typecho_contents WHERE cid=? AND type='page'`, cid)
	return err
}

// GetPosts  根据条件查询多条文章 状态 条数 页数
func (s *QPU) GetPosts(status string, limit, pageNum int) error {
	err := db.Select(&s.PostArr, `
		SELECT * FROM  typecho_contents
		WHERE type='post' AND status=?
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, status, limit, pageNum*limit-limit)
	return err
}

// GetPage  根据条件查询单独页面
func (s *QPU) GetPage(p string) error {
	err := db.Select(&s.PageArr, `
		SELECT * FROM  typecho_contents 
		WHERE type='page' AND slug = ?`, p)
	if len(s.PageArr) == 0 {
		return errors.New("not found")
	}
	return err
}

// GetPages  根据条件查询多条页面
func (s *QPU) GetPages() error {
	err := db.Select(&s.PageArr, `
		SELECT * FROM  typecho_contents 
		WHERE type='page'
		ORDER BY "order" `)
	return err
}

// CheckNext 查询是否有下一页，如果有则ture
func (s *QPU) CheckNext() error {
	err := db.Get(s.HaveNext, `
		SELECT cid FROM  typecho_contents 
		WHERE cid>?`, s.PostArr[len(s.PostArr)-1].Cid)
	return err
}
