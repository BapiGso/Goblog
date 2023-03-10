package moe

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

// Hash 计算字符串sha1
func Hash(input string) string {
	h := sha1.New() // md5加密类似md5.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(input))
	bs := h.Sum(nil)
	h.Reset()
	passwdhash := hex.EncodeToString(bs)
	return passwdhash
}

// IsNum 首页返回1，不是数字返回err调用404，其他为对应页数
func IsNum(numstr string) (uint64, error) {
	if numstr == "" {
		return 1, nil
	}
	num, err := strconv.ParseUint(numstr, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func TimeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("time cost = %v\n", tc)
	}
}

func IsLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//登录页面不用这个中间件
		if c.Path() == "/admin" {
			return next(c)
		}
		//后台页面没有cookie的全部跳去登录
		sess, _ := session.Get("smoesession", c)
		if sess.Values["isLogin"] != true {
			return c.Redirect(http.StatusFound, "/admin")
		}
		return next(c)
	}
}
