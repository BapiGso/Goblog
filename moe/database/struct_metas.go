package database

type Metas struct {
	Mid         int    `db:"mid"`
	Name        string `db:"name"`
	Slug        string `db:"slug"`
	Type        string `db:"type"`
	Description string `db:"description"`
	Count       int    `db:"count"`
	Order       int    `db:"order"`
	Parent      int    `db:"parent"`
}
