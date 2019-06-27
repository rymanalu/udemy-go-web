package main

import (
	"html/template"
	"log"
	"net/http"
)

type me struct {
	Name string
	Age  int
}

type indexHandler int

func (h indexHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	handleError(tpl.ExecuteTemplate(res, "index.gohtml", nil))
}

type dogHandler int

func (h dogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	handleError(tpl.ExecuteTemplate(res, "dog.gohtml", nil))
}

type meHandler int

func (h meHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	i := me{"Roni Yusuf", 24}

	handleError(tpl.ExecuteTemplate(res, "me.gohtml", i))
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.Handle("/", new(indexHandler))

	http.Handle("/dog/", new(dogHandler))

	http.Handle("/me/", new(meHandler))

	http.ListenAndServe(":8080", nil)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
