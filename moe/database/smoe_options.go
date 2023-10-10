package database

// GetOption
func (s *QPU) GetOption(name string) (string, error) {
	var value string
	err := db.Get(&value, `
		SELECT * FROM  typecho_options 
		WHERE name=?`, name)
	return value, err
}
