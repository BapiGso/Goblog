package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func AccessArr(Db *sqlx.DB, limit, pageNum int) []Access {
	var data []Access
	err := Db.Select(&data, `SELECT * FROM  main.typecho_access_log
        WHERE robot=0
		ORDER BY id DESC 
		LIMIT ? OFFSET ?`, limit, pageNum*limit-limit)
	fmt.Println(err)
	return data
}
