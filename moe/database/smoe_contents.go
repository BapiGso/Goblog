package database

import (
	"errors"
	"github.com/labstack/gommon/log"
)

// Media 查询文件组，后台专用
func (s *S) Media(limit, pageNum int) error {
	err := db.Select(&s.PostArr, `
		SELECT * FROM  typecho_contents
		WHERE type='attachment'
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, limit, pageNum*limit-limit)
	return err
}

// GetPostWithCid 根据Cid查询文章
func (s *S) GetPostWithCid(cid int) error {
	err := db.Select(&s.PostArr, `SELECT * FROM typecho_contents WHERE cid=? AND type='post'`, cid)
	if len(s.PostArr) == 0 {
		return errors.New("没有查询到结果")
	}
	return err
}

// GetPageWithCid 根据Cid查询独立页面
func (s *S) GetPageWithCid(cid int) error {
	err := db.Select(&s.PageArr, `SELECT * FROM typecho_contents WHERE cid=? AND type='page'`, cid)
	return err
}

// GetPosts  根据条件查询多条文章 状态 条数 页数
func (s *S) GetPosts(status string, limit, pageNum int) error {
	err := db.Select(&s.PostArr, `
		SELECT * FROM  typecho_contents
		WHERE type='post' AND status=?
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, status, limit, pageNum*limit-limit)
	return err
}

// GetPages  根据条件查询多条页面
func (s *S) GetPages() error {
	err := db.Select(&s.PageArr, `
		SELECT * FROM  typecho_contents 
		WHERE type='page'
		ORDER BY 'order' `)
	return err
}

// HaveNext 查询是否有下一页，如果有则ture
func (s *S) HaveNext() bool {
	var data int
	err := db.Get(data, `
		SELECT cid FROM  typecho_contents 
		WHERE cid>?`)
	if err != nil {
		log.Error(err)
	}
	return true
}
