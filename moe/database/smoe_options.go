package database

// GetOption 1
func (q *QPU) GetOption(name string) (string, error) {
	var value string
	err := db.Get(&value, `
		SELECT value FROM  typecho_options 
		WHERE name=?`, name)
	return value, err
}
