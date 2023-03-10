package moe

import (
	"SMOE/moe/archive"
	"SMOE/moe/bangumi"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (s *Smoe) Page(c echo.Context) error {
	data := 0
	return c.Render(200, "page.template", data)
}

func (s *Smoe) Archive(c echo.Context) error {
	data := archive.QueryTime(s.Db)
	return c.Render(200, "page-timeline.template", data)
}

// Bangumi todo https://freefrontend.com/css-cards/
func (s *Smoe) Bangumi(c echo.Context) error {
	timeUnix := time.Now().Unix()
	if timeUnix-bangumi.Bgmcache.TTL > 604800 {
		url := bangumi.QueryBGM(s.Db)
		bangumi.CurlBGM(url)
	}
	return c.Render(http.StatusOK, "page-bangumi.template", bangumi.Bgmcache)
}
