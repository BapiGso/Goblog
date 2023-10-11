package database

import (
	"fmt"
	"math/big"
	"net"
	"strings"
	"time"
)

type Access struct {
	ID               int     `db:"id"`
	UA               string  `db:"ua"`
	BrowserID        string  `db:"browser_id"`
	BrowserVersion   string  `db:"browser_version"`
	OSID             string  `db:"os_id"`
	OSVersion        string  `db:"os_version"`
	URL              string  `db:"url"`
	Path             string  `db:"path"`
	QueryString      *string `db:"query_string"`
	IP               int64   `db:"ip"`
	Entrypoint       string  `db:"entrypoint"`
	EntrypointDomain *string `db:"entrypoint_domain"`
	Referer          string  `db:"referer"`
	RefererDomain    *string `db:"referer_domain"`
	Time             int64   `db:"time"`
	ContentID        *int64  `db:"content_id"`
	MetaID           *int64  `db:"meta_id"`
	Robot            int8    `db:"robot"`
	RobotID          string  `db:"robot_id"`
	RobotVersion     string  `db:"robot_version"`
}

// IPString  数字转ip地址
func (a Access) IPString() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(a.IP>>24), byte(a.IP>>16), byte(a.IP>>8), byte(a.IP))
}

// IPInt  ip地址转数字
func IPInt(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

func (a Access) UnixFormat() string {
	format := time.Unix(a.Time, 0).Format("2006-01-02 15:04:05")
	return format
}

// SimplifyUserAgent 简化浏览器的User-Agent字符串并检测爬虫（bot）
func (a Access) SimplifyUserAgent() string {
	browsers := []string{"Chrome", "Safari", "Firefox", "Edge", "Opera", "MSIE", "Trident"}
	bots := []string{"Googlebot", "Bingbot", "Slurp", "DuckDuckBot", "Baiduspider", "YandexBot", "Sogou", "Exabot", "facebot", "ia_archiver"}

	for _, bot := range bots {
		if strings.Contains(a.UA, bot) {
			return bot
		}
	}

	for _, browser := range browsers {
		if strings.Contains(a.UA, browser) {
			return browser
		}
	}

	return "Unknown"
}

func (a Access) SubReferer() string {
	// 将字符串转换为[]rune，以便正确处理Unicode字符
	runes := []rune(a.Referer)
	runesLength := len(runes)

	if runesLength <= 15 {
		return a.Referer
	}

	more := fmt.Sprintf(`...<a class="tooltip" data-tooltip="%v">查看更多</a>`, a.Referer)
	return string(runes[:15]) + more
}
