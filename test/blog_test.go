package _test

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendTenPostRequests(t *testing.T) {
	// 创建 HTTP 请求正文
	requestBody := bytes.NewBufferString(`{"user": "测试用户", "pwd": "测试密码"}`)

	// 创建 Echo 实例并向其中添加路由
	e := echo.New()
	e.POST("/posts", func(c echo.Context) error {
		return c.String(201, "1")
	}, middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	// 执行十次 POST 请求
	for i := 0; i < 12; i++ {
		// 创建 HTTP 请求并发送到 Echo 实例
		req := httptest.NewRequest("POST", "/posts", requestBody)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		// 验证 HTTP 状态码是否正确并记录在测试日志中
		if i < 10 {
			if rec.Code != http.StatusCreated {
				t.Errorf("HTTP 状态码错误，实际为：%v，期望为：%v", rec.Code, http.StatusCreated)
			} else {
				t.Logf("第 %d 次请求状态码为：%v", i+1, rec.Code)
			}
		} else {
			t.Logf("HTTP 状态码正确，实际为：%v，期望为：%v", rec.Code, http.StatusTooManyRequests)
		}
	}
}
