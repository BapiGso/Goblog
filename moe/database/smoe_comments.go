package database

import (
	"errors"
	"time"
)

// GetComms  根据条件查询多条评论 状态 条数 页数
func (q *QPU) GetComms(status string, limit, pageNum int) error {
	err := DB.Select(&q.Comments, `
		SELECT * FROM  typecho_comments
		WHERE status=?
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, status, limit+1, pageNum*limit-limit)
	//多查一个判断是否有下一页
	if len(q.Comments) == limit+1 {
		//q.HaveNext = pageNum + 1
		q.Contents = q.Contents[:len(q.Contents)-1]
	}
	return err
}

// CommentsWithCid 根据文章cid查询该文章的评论
func (q *QPU) CommentsWithCid(status string, cid int) error {
	err := DB.Select(&q.Comments, `SELECT * FROM  typecho_comments 
		WHERE cid=?
		AND status=?
		ORDER BY created`, cid, status)
	return err
}

// SortComments 文章页面的评论区用到了该函数 排序评论
func (q *QPU) SortComments() [][]Comments {
	var final [][]Comments
	parentMap := make(map[uint]int)
	for _, v := range q.Comments {
		//父评论新建一个组，因为按时间排序肯定比子评论先
		if v.Parent == 0 {
			//初始化tmp的同时就把v加入切片
			tmp := []Comments{v}
			final = append(final, tmp)
			parentMap[v.Coid] = len(final) - 1
		} else { //如果是子评论，找自己属于哪个父评论组
			index := parentMap[v.Parent]
			final[index] = append(final[index], v)
			parentMap[v.Coid] = index
		}
	}
	return final
}

// InsertComment 已完成
func InsertComment(data map[string]any) error {
	insertData := Comments{
		Coid:     0,
		Cid:      data["Cid"].(uint),
		Parent:   data["Parent"].(uint),
		Created:  time.Now().Unix(),
		Author:   data["Author"].(string),
		Mail:     data["Mail"].(string),
		Ip:       data["Ip"].(string),
		Agent:    data["Agent"].(string),
		Text:     data["Text"].(string),
		Status:   "waiting",
		AuthorId: 0,
		Url: func(url string) *string {
			if url == "" {
				return nil
			}
			return &url
		}(data["Url"].(string)),
	}
	tx, err := DB.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}
	var count int
	if err = tx.Get(&count, "SELECT COUNT(*) FROM typecho_contents WHERE cid = ?", data["Cid"]); err != nil {
		tx.Rollback()
		return err
	}
	if count == 0 {
		tx.Rollback()
		return errors.New("duplicate cid not found in table contents")
	}
	if err = tx.Get(&count, "SELECT COUNT(*) FROM typecho_comments WHERE coid = ?", data["Parent"]); err != nil {
		tx.Rollback()
		return err
	}
	//如果表单有写父评论但是却查不到父评论
	if data["Parent"].(uint) != 0 && count == 0 {
		tx.Rollback()
		return errors.New("duplicate coid not found in table comments")
	}
	if err = tx.Get(&insertData.Coid, `SELECT seq+1 FROM sqlite_sequence WHERE name='typecho_comments'`); err != nil {
		tx.Rollback()
		return err
	}
	if _, err = tx.NamedExec(`INSERT INTO typecho_comments 
		VALUES (:coid,:cid,:created,:author,:authorId,:ownerId,:mail,:url,:ip,:agent,:text,:type,:status,:parent)`, insertData); err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
