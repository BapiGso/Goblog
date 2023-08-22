package database

func (s *S) UserWithName(name string) error {
	err := db.Get(&s.UserInfo, `SELECT * FROM  typecho_users WHERE name = ?`, name)
	return err
}
