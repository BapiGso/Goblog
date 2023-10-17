package database

import (
	"errors"
	"time"
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

func UpdateView(cid string) error {
	_, err := db.Exec(`UPDATE typecho_contents SET views = views + 1 WHERE cid = ?`, cid)
	if err != nil {
		return err
	}
	return nil
}

func (q *QPU) SortContents() any {
	type yData struct {
		Mon   string
		MData []Contents
	}
	type Data struct {
		Year  string
		YData []yData
	}
	var final []Data
	for _, v := range q.PostArr {
		year := time.Unix(v.Created, 0).Format("2006")
		mon := time.Unix(v.Created, 0).Format("01")
		if len(final) != 0 && final[len(final)-1].Year == year {
			if final[len(final)-1].YData[len(final[len(final)-1].YData)-1].Mon == mon {
				final[len(final)-1].YData[len(final[len(final)-1].YData)-1].MData = append(final[len(final)-1].YData[len(final[len(final)-1].YData)-1].MData, v)
			} else {
				final[len(final)-1].YData = append(final[len(final)-1].YData, yData{
					Mon:   mon,
					MData: []Contents{v},
				})
			}
		} else {
			final = append(final, Data{
				year,
				[]yData{
					{
						mon,
						[]Contents{v},
					},
				},
			},
			)
		}
	}
	return final
}
