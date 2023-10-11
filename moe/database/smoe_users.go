package database

func (q *QPU) UserWithName(name string) error {
	err := db.Get(&q.UserInfo, `SELECT * FROM  typecho_users WHERE name = ?`, name)
	return err
}
