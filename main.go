package main

import (
	_ "github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"io"
	"main/smoe"
	_ "main/smoe"
	_ "modernc.org/sqlite"
	_ "net/http/pprof"
	"text/template"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	s := Smoe.New()
	s.BindFlag()
	s.InitializeDatabase()
	s.Middleware()
	s.LoadRoutes()
	s.Listen()
}
