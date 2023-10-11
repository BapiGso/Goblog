package database

// CommentsWithCid 根据文章cid查询该文章的评论
func (q *QPU) CommentsWithCid(cid int) error {
	err := db.Select(&q.CommArr, `SELECT * FROM  typecho_comments 
		WHERE cid=?
		ORDER BY created`, cid)
	return err
}

// SortComments 排序评论
func (q *QPU) SortComments() [][]Comments {
	var final [][]Comments
	parentMap := make(map[uint]int)
	for _, v := range q.CommArr {
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

// InsertComment todo
func InsertComment(data map[string]any) error {
	arg := Comments{}
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	_, err = tx.NamedExec(`INSERT `, arg)
	if err != nil {
		return err
	}
	return nil
}
