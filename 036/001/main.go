package main

import (
	"io"
	"net/http"
)

type handler int

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		io.WriteString(w, "index")
	case "/customers":
		io.WriteString(w, "customer customer customer")
	case "/admins":
		io.WriteString(w, "admin admin admin")
	default:
		io.WriteString(w, "fallback")
	}
}

func main() {
	h := new(handler)

	http.ListenAndServe(":8080", h)
}
