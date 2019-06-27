package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/about", about)
	mux.GET("/users/:id", users)

	http.ListenAndServe(":8080", mux)
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	handleError(err)
}

func about(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "about.gohtml", nil)
	handleError(err)
}

func users(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "users.gohtml", ps.ByName("id"))
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
