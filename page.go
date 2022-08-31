package main

import "github.com/labstack/echo/v4"

type PageData struct {
	Title     string
	Text      []byte
	Text2HTML string
}

func queryPage(slug string) PageData {
	data := PageData{}
	_ = db.QueryRow(`SELECT title,text
		FROM typecho_contents
		WHERE slug=?`, slug).Scan(&data.Title, &data.Text)
	data.Text2HTML = md2html(data.Text)
	return data
}

func Page(c echo.Context) error {
	slug := c.Param("page")
	data := queryPage(slug)
	return c.Render(200, "page.template", data)
}
