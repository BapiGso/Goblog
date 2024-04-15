package mymiddleware

import (
	"bytes"
	"encoding/binary"
	"github.com/coocood/freecache"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hash/fnv"
	"io"
	"net/http"
)

type CacheConfig struct {
	Skipper func(c echo.Context) bool
	Expire  int //缓存过期时间,秒
}

var DefaultCacheConfig = CacheConfig{
	Skipper: middleware.DefaultSkipper,
	Expire:  60,
}

func Cache() echo.MiddlewareFunc {
	return CacheWithConfig(DefaultCacheConfig)
}

type responseRetainer struct {
	io.Writer
	http.ResponseWriter
}

// CacheWithConfig 可以看看这篇文章的实现 https://medium.com/@anajankow/speed-up-your-get-requests-with-a-cache-middleware-e3584f2e4ef1
// todo 还没写完，后面再写吧，懒得弄了
func CacheWithConfig(config CacheConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		//config.Skipper = DefaultBrotliConfig.Skipper
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}
			//如果不为GET请求则不用这个中间件
			if c.Request().Method != http.MethodGet {
				return next(c)
			}
			reqHeader, resBody := new(bytes.Buffer), new(bytes.Buffer)
			if err := binary.Write(reqHeader, binary.LittleEndian, c.Request().Header); err != nil {
				return err
			}
			mw := io.MultiWriter(c.Response().Writer, resBody)
			_ = &responseRetainer{Writer: mw, ResponseWriter: c.Response().Writer}
			//c.Response().Writer = writer

			if err := next(c); err != nil {
				c.Error(err)
			}
			//检查是否有缓存，如果有则用缓存来响应请求
			if _, err := store.Get(fnv.New32a().Sum(reqHeader.Bytes())); err != nil {
				//如果没有缓存，响应结束后将渲染用到的数据写入缓存
				c.Response().After(func() {

				})
				return next(c)
			} else {
				//如果有缓存，则用缓存来响应
				return next(c)
			}

		}
	}
}

// 申请一个100M大小的缓存
var store = freecache.NewCache(100 * 1024 * 1024)
