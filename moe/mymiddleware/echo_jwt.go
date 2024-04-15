package mymiddleware

import (
	"flag"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// JWTKey 生成一个随机[]byte
var JWTKey = []byte(strconv.Itoa(rand.Int()))

func init() {
	debug := flag.Bool("debug", false, "enable debug mode")
	// 解析传入的命令行参数
	flag.Parse()
	if *debug {
		JWTKey = []byte("123")
	}
}

func JWT() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			restricted := strings.HasPrefix(c.Path(), "/admin/") //判断当前路径是否是受限制路径（除了登录页面以外的后台路径）
			_, err := c.Cookie("smoe_token")
			return err != nil && !restricted //如果读不到cookie且不是受限制路径就跳过
		},
		ErrorHandler: func(c echo.Context, err error) error {
			//todo 触发错误后ip限制
			c.SetCookie(&http.Cookie{Name: "smoe_token", Expires: time.Now(), MaxAge: -1, HttpOnly: true})
			return echo.ErrTeapot
		},
		SuccessHandler: func(c echo.Context) {

		},
		SigningKey:  JWTKey,
		TokenLookup: "cookie:smoe_token",
	})
}
