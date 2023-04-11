package test

import (
	"bytes"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// 测试首页是否正常
func TestIndex(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/", nil)
	rec := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rec, req)
	t.Logf("HTTP 状态码，实际为：%v，期望为：%v，请检查应用是否已运行在80端口", rec.Code, http.StatusOK)
}

// 测试管理页权限是否正常
func TestAdminWithCookie(t *testing.T) {
	// 随机生成一个 cookie
	userID := strconv.Itoa(rand.Int())
	cookie := &http.Cookie{Name: "userId", Value: userID}

	req := httptest.NewRequest("GET", "http://localhost/admin/test", nil)
	req.AddCookie(cookie)

	rec := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rec, req)

	t.Logf("HTTP 状态码，实际为：%v，期望为：%v", rec.Code, http.StatusFound)
}

// 测试静态文件是否正常
func TestAssets(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/favicon.ico", nil)
	rec := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rec, req)
	t.Logf("HTTP 状态码，实际为：%v，期望为：%v", rec.Code, http.StatusOK)
}

// 测试登录页防爆破
func TestSendTenPostRequests(t *testing.T) {
	// 创建 HTTP 请求正文
	requestBody := bytes.NewBufferString(`{"user": "测试用户", "pwd": "测试密码"}`)
	// 执行十次 POST 请求
	for i := 0; i < 11; i++ {
		// 创建 HTTP 请求并发送到 Echo 实例
		req := httptest.NewRequest("POST", "http://localhost/admin", requestBody)
		rec := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rec, req)
		if i >= 10 {
			t.Logf("HTTP 状态码，实际为：%v，期望为：%v", rec.Code, http.StatusTooManyRequests)
		}
	}
}

// 测试错误的文章
func TestPost(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/archive/abc", nil)
	rec := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rec, req)
	t.Logf("HTTP 状态码，实际为：%v，期望为：%v", rec.Code, http.StatusNotFound)
}

// 测试错误的页面
func TestPage(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/page/abc", nil)
	rec := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rec, req)
	t.Logf("HTTP 状态码，实际为：%v，期望为：%v", rec.Code, http.StatusNotFound)
}
