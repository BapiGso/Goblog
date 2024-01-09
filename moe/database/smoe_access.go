package database

func InsertAccess(data map[string]any) error {
	err := DB.Select(&data, `SELECT * FROM  main.typecho_access_log
        WHERE robot=0
		ORDER BY id DESC
		LIMIT ? OFFSET ?`)
	return err
}

// Access  查询日志
func (q *QPU) GetAccess(limit, pageNum int) error {
	err := DB.Select(&q.Access, `
		SELECT * FROM  typecho_access_log
		ORDER BY ROWID DESC
		LIMIT ? OFFSET ?`, limit, pageNum*limit-limit)
	return err
}
