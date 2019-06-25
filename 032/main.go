package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type httpHandler int

func (h httpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	h := new(httpHandler)

	http.ListenAndServe(":8080", h)
}
