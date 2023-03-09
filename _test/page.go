package _test

import "github.com/labstack/echo/v4"

type PageData struct {
	Cid       int
	Title     string
	Text      []byte
	Text2HTML string
}

func queryPage(slug string) PageData {
	data := PageData{}
	//_ = db.QueryRow(`SELECT cid,title,text
	//	FROM typecho_contents
	//	WHERE slug=?`, slug).Scan(&data.Cid, &data.Title, &data.Text)
	return data
}

func Page(c echo.Context) error {
	slug := c.Param("page")
	data := queryPage(slug)
	if data.Cid == 0 {
		return echo.ErrNotFound
	}
	return c.Render(200, "page.template", data)
}
