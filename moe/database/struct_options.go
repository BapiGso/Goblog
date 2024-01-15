package database

type Options struct {
	Name  string `db:"name"`
	User  string `db:"user"`
	Value string `db:"value"`
}
