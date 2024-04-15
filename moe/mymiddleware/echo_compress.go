package mymiddleware

import (
	"github.com/CAFxX/httpcompression"
	"github.com/klauspost/compress/gzip"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// AutoCompression 自动响应压缩，支持zstd，brotli，gzip
var adapter = func() echo.MiddlewareFunc {
	a, err := httpcompression.DefaultAdapter(httpcompression.GzipCompressionLevel(gzip.DefaultCompression))
	if err != nil {
		return middleware.Gzip()
	}

	return middleware.Gzip()
	return echo.WrapMiddleware(a)
}()

func AutoCompression() echo.MiddlewareFunc {
	return adapter
}
