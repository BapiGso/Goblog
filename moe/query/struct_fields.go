package query

type Fields struct {
	Cid        int     `db:"cid"`
	Name       string  `db:"name"`
	Type       string  `db:"type"`
	StrValue   string  `db:"str_value"`
	IntValue   int     `db:"int_value"`
	FloatValue float64 `db:"float_value"`
}
