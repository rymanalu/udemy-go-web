package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(res, "Do my search: "+v)
}

// Visit this page:
// http://localhost:8080/?q=the%20beatles
