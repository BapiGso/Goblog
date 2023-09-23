package database

import (
	"database/sql"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	_ "modernc.org/sqlite"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	var err error
	if _, err := os.Stat("../../usr/smoe.db"); os.IsNotExist(err) {
		// 文件不存在
		log.Fatal("数据库文件不存在，请检查路径是否正确")
	}
	db, err = sqlx.Connect("sqlite", "../../usr/smoe.db")
	if err != nil {
		result := m.Run() //运行go的测试，相当于调用main方法
		os.Exit(result)   //退出程序
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		qpu := &QPU{}
		json.Unmarshal(nil, qpu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		qpu := NewQPU()
		json.Unmarshal(nil, qpu)
		qpu.Free()
	}
}

func BenchmarkS_GetPosts(b *testing.B) {
	qpu := NewQPU()
	for i := 0; i < b.N; i++ {
		qpu.GetPosts("publish", 10, 1)
	}
}

func Benchmark_NaiveSql(b *testing.B) {
	db2, _ := sql.Open("sqlite", "../../usr/smoe.db")
	for i := 0; i < b.N; i++ {
		rows, _ := db2.Query(`SELECT * FROM typecho_contents WHERE status='publish' ORDER BY cid LIMIT 10 `)
		defer rows.Close()
		var fos []Contents
		for rows.Next() {
			f := Contents{}
			rows.Scan(&f.Cid, &f.Title, &f.Slug, &f.Created, &f.Modified, &f.Text, &f.Order,
				&f.AuthorId, &f.Template, &f.Type, &f.Status, &f.Password, &f.CommentsNum,
				&f.AllowComment, &f.AllowPing, &f.AllowFeed, &f.Parent, &f.Views, &f.Likes)
			fos = append(fos, f)
		}
	}
}

func BenchmarkS_GetPostWithCid(b *testing.B) {
	qpu := NewQPU()
	for i := 0; i < b.N; i++ {
		qpu.GetPostWithCid(166)
	}
}
