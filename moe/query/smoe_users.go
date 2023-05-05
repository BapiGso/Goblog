package query

import "github.com/jmoiron/sqlx"

func UserWithName(Db *sqlx.DB, name string) (User, error) {
	var data User
	err := Db.Get(&data, `SELECT * FROM  typecho_users WHERE name = ?`, name)
	return data, err
}
