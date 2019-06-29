package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/pics/", http.StripPrefix("/pics", http.FileServer(http.Dir("./public/pics"))))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.gohtml"))

	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)

	if err != nil {
		log.Fatalln(err)
	}
}
