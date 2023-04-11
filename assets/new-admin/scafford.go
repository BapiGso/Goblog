package main

import (
	"net/http"
	"text/template"
)

var temp, err = template.ParseGlob("new-admin/*.template")

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.Handle("/new-admin/", http.FileServer(http.Dir("")))
	http.ListenAndServe(":8848", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "test.template", nil)
}
