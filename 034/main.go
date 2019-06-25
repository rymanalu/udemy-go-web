package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type httpHandler int

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Manalu-Key", "This is from Manalu")

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	h := new(httpHandler)

	http.ListenAndServe(":8080", h)
}
