package main

import (
	"html/template"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/cat/", cat)
	http.HandleFunc("/cat.jpeg", catImage)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(res, "foo ran")

	handleError(err, res)
}

func cat(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("cat.gohtml"))

	err := tpl.ExecuteTemplate(res, "cat.gohtml", nil)

	handleError(err, res)
}

func catImage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "cat.jpeg")
}

func handleError(err error, res http.ResponseWriter) {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
