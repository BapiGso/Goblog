package moe

import (
	"fmt"
	"github.com/BapiGso/SMOE/moe/query"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func isLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//登录页面不用这个中间件
		if c.Path() == "/admin" {
			return next(c)
		}
		//后台页面没有cookie的全部跳去登录
		sess, _ := session.Get("smoeSession", c)
		if sess.Values["isLogin"] != true {
			return c.Redirect(http.StatusFound, "/admin")
		}
		return next(c)
	}
}

func attachDB(db *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}

func SetDefaultQueryParams(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		queryParams := map[string]string{
			"commstatus": "approved",
			"status":     "publish",
			"page":       "1",
		}
		for key, value := range queryParams {
			if c.QueryParam(key) == "" {
				c.QueryParams().Set(key, value)
			}
		}
		return next(c)
	}
}

func accessLogMiddleware(db *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 记录开始时间
			//startTime := time.Now()

			// 继续处理请求
			err := next(c)
			if err != nil {
				c.Error(err)
			}

			// 记录结束时间
			//endTime := time.Now()

			// 收集信息
			req := c.Request()
			//ip := c.RealIP()
			ua := req.UserAgent()
			url := req.URL
			path := req.URL.Path
			//query1 := req.URL.RawQuery
			referer := req.Referer()

			if len(referer) > 0 {
				//refURL, err := url.Parse(referer)
				//if err == nil {
				//	refererDomain = refURL.Hostname()
				//}
			}

			entrypoint := url
			//entrypointDomain := req.Host
			//duration := endTime.Sub(startTime)
			//time := duration.Milliseconds()

			// 插入数据库
			logEntry := query.Access{
				UA:   ua,
				URL:  url.String(),
				Path: path,

				//IP:               ip,
				Entrypoint: entrypoint.String(),
				//EntrypointDomain: entrypointDomain,
				Referer: referer,
				//RefererDomain: refererDomain,
				//Time:             int32(time),
			}
			_, err = db.NamedExec(`INSERT INTO typecho_access_log (ua, url, path, query_string, ip, entrypoint, entrypoint_domain, referer, referer_domain, time) VALUES (:ua, :url, :path, :query_string, :ip, :entrypoint, :entrypoint_domain, :referer, :referer_domain, :time)`, logEntry)
			if err != nil {
				fmt.Printf("Error inserting log entry: %v\n", err)
			}

			return err
		}
	}
}
