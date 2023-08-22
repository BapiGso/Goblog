package database

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if _, err := os.Stat("../../usr/smoe.db"); os.IsNotExist(err) {
		// 文件不存在
		log.Fatal("数据库文件不存在，请检查路径是否正确")
	}
	db, _ = sqlx.Connect("sqlite", "../../usr/smoe.db")
	result := m.Run() //运行go的测试，相当于调用main方法
	os.Exit(result)   //退出程序
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		qpu := &S{}
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

}

func BenchmarkS_GetPostWithCid(b *testing.B) {
	qpu := NewQPU()
	for i := 0; i < b.N; i++ {
		qpu.GetPostWithCid(166)
	}
}
