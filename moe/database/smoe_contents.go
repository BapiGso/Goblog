package database

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

// Deprecated: 查询文件组，后台专用,不查文件了，没啥用
func (q *QPU) Media(limit, pageNum int) error {
	err := DB.Select(&q.Contents, `
		SELECT * FROM  typecho_contents
		WHERE type='attachment'
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, limit, pageNum*limit-limit)
	return err
}

// GetPostWithCid 根据Cid查询文章 todo 权限
func (q *QPU) GetPostWithCid(status string, cid int) error {
	err := DB.Select(&q.Contents, `
		SELECT * FROM typecho_contents 
		WHERE cid=? AND status=?
		AND type='post'`, cid, status)
	if len(q.Contents) == 0 {
		return errors.New("not found")
	}
	return err
}

// GetWithCid 根据Cid查询文章或页面，写文章用的
func (q *QPU) GetWithCid(cid int) error {
	err := DB.Select(&q.Contents, `
		SELECT * FROM typecho_contents 
        WHERE cid=?`, cid)
	return err
}

// GetPosts  根据条件查询多条文章 状态 条数 页数
func (q *QPU) GetPosts(status string, limit, pageNum int) error {
	err := DB.Select(&q.Contents, `
		SELECT * FROM  typecho_contents
		WHERE type='post' AND status=?
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, status, limit+1, pageNum*limit-limit)
	//多查一个判断是否有下一页
	if len(q.Contents) == limit+1 {
		//q.HaveNext = pageNum + 1
		q.Contents = q.Contents[:len(q.Contents)-1]
	}
	return err
}

// GetPage  根据条件查询单独页面
func (q *QPU) GetPage(p string) error {
	err := DB.Select(&q.Contents, `
		SELECT * FROM  typecho_contents 
		WHERE type='page' AND slug = ?`, p)
	if len(q.Contents) == 0 {
		return errors.New("not found")
	}
	return err
}

// GetPages  根据条件查询多条页面
func (q *QPU) GetPages() error {
	err := DB.Select(&q.Contents, `
		SELECT * FROM  typecho_contents 
		WHERE type='page'
		ORDER BY "order" `)
	return err
}

func UpdateView(cid string) error {
	_, err := DB.Exec(`UPDATE typecho_contents SET views = views + 1 WHERE cid = ?`, cid)
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: 已弃用 (archives页面排列时间线调用了这个函数
func (q *QPU) SortTimeline() any {
	type yData struct {
		Mon   string
		MData []Contents
	}
	type Data struct {
		Year  string
		YData []yData
	}
	var final []Data
	for _, v := range q.Contents {
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

// InsertContent 插入新的文章
func InsertContent(data map[string]any) error {
	insertData := Contents{
		Cid:          data["Cid"].(int),
		Title:        data["Title"].(string),
		Slug:         data["Slug"].(string),
		Created:      time.Now().Unix(),
		Modified:     time.Now().Unix(),
		Text:         []byte(data["Text"].(string)),
		Order:        0,
		AuthorId:     1,
		Template:     nil,
		Type:         data["Type"].(string),
		Status:       "publish",
		Password:     nil,
		AllowComment: 1,
		AllowPing:    0,
		AllowFeed:    0,
		CommentsNum:  0,
		Parent:       0,
		Views:        0,
		Likes:        0,
	}
	tx, err := DB.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Get(&insertData.Cid, `SELECT seq FROM sqlite_sequence WHERE name='typecho_contents'`); err != nil {
		tx.Rollback()
		return err
	}
	n, _ := rand.Int(rand.Reader, big.NewInt(10))
	insertData.Cid = insertData.Cid + int(n.Int64()) + 1
	if insertData.Slug == "" {
		insertData.Slug = strconv.Itoa(insertData.Cid)
	}
	if _, err = tx.NamedExec(`INSERT INTO typecho_contents 
		VALUES (:cid,:title,:slug,:created,:modified,:text,:order,
		        :authorId,:template,:type,:status,:password,:commentsNum,
		        :allowComment,:allowPing,:allowFeed,:parent,:views,:likes)`, insertData); err != nil {
		tx.Rollback()
		return err
	}
	if _, err = tx.Exec(`INSERT INTO typecho_fields 
		VALUES (?,'coverList','str',?,0,0.0)`, insertData.Cid, data["CoverList"].(string)); err != nil {
		tx.Rollback()
		return err
	}
	if _, err = tx.Exec(`INSERT INTO typecho_fields 
		VALUES (?,'musicList','str',?,0,0.0)`, insertData.Cid, data["MusicList"].(string)); err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// UpdateContent 更新文章 todo
func UpdateContent(data map[string]any) error {
	fmt.Println(data)
	insertData := Contents{
		Cid:      data["Cid"].(int),
		Title:    data["Title"].(string),
		Slug:     data["Slug"].(string),
		Modified: time.Now().Unix(),
		Text:     []byte(data["Text"].(string)),
		Order:    0,
	}
	tx, err := DB.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err = tx.NamedExec(`UPDATE typecho_contents
		SET title=:title,slug=:slug,modified=:modified,text=:text,"order"=:order
		WHERE cid=:cid`, insertData); err != nil {
		tx.Rollback()
		return err
	}
	if _, err = tx.Exec(`UPDATE typecho_fields
		SET str_value=? 
		WHERE cid=? AND name='coverList'`, data["CoverList"].(string), insertData.Cid); err != nil {
		tx.Rollback()
		return err
	}
	if _, err = tx.Exec(`UPDATE typecho_fields
		SET str_value=? 
		WHERE cid=? AND name='musicList'`, data["MusicList"].(string), insertData.Cid); err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
