package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Your request method at index: %s\n\n", req.Method)
}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Your request method at foo: %s\n\n", req.Method)

	// Can be this way...
	// res.Header().Set("Location", "/")
	// res.WriteHeader(http.StatusSeeOther)

	// Or this way...
	http.Redirect(res, req, "/", http.StatusMovedPermanently)

	// Statuses:
	// http.StatusSeeOther
	// http.StatusTemporaryRedirect
	// http.StatusMovedPermanently
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Your request method at bar: %s\n\n", req.Method)
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
