package query

import (
	"fmt"
	"math/big"
	"net"
)

type Access struct {
	ID               int    `db:"id"`
	UA               string `db:"ua"`
	BrowserID        string `db:"browser_id"`
	BrowserVersion   string `db:"browser_version"`
	OSID             string `db:"os_id"`
	OSVersion        string `db:"os_version"`
	URL              string `db:"url"`
	Path             string `db:"path"`
	QueryString      string `db:"query_string"`
	IP               int64  `db:"ip"`
	Entrypoint       string `db:"entrypoint"`
	EntrypointDomain string `db:"entrypoint_domain"`
	Referer          string `db:"referer"`
	RefererDomain    string `db:"referer_domain"`
	Time             int64  `db:"time"`
	ContentID        int    `db:"content_id"`
	MetaID           int    `db:"meta_id"`
	Robot            int8   `db:"robot"`
	RobotID          string `db:"robot_id"`
	RobotVersion     string `db:"robot_version"`
}

// IPString  数字转ip地址
func (a *Access) IPString() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(a.IP>>24), byte(a.IP>>16), byte(a.IP>>8), byte(a.IP))
}

// IPInt  ip地址转数字
func IPInt(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}
