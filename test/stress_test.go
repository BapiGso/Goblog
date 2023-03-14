package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkEchoHandler(b *testing.B) {
	// 创建请求
	req := httptest.NewRequest("GET", "http://localost/", nil)

	// 循环执行请求
	for i := 0; i < b.N; i++ {
		// 发送请求并获取响应
		rec := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rec, req)

		// 检查响应码是否为 200
		if rec.Code != http.StatusOK {
			b.Fatalf("响应码不正确：%d", rec.Code)
		}
	}
}
