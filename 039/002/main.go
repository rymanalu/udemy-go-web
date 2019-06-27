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

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		handleError(tpl.ExecuteTemplate(res, "index.gohtml", nil))
	})

	http.HandleFunc("/dog/", func(res http.ResponseWriter, req *http.Request) {
		handleError(tpl.ExecuteTemplate(res, "dog.gohtml", nil))
	})

	http.HandleFunc("/me/", func(res http.ResponseWriter, req *http.Request) {
		i := me{"Roni Yusuf", 24}

		handleError(tpl.ExecuteTemplate(res, "me.gohtml", i))
	})

	http.ListenAndServe(":8080", nil)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
