package mymiddleware

import (
	"github.com/labstack/echo/v4"
	"io"
	"text/template"
)

type TemplateRender struct {
	Template *template.Template //渲染模板
}

// Render todo 可以在这里检查data是否有cache，有的话用cache，没有的话添加进cache
func (t *TemplateRender) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.Template.ExecuteTemplate(w, name, data)
}
