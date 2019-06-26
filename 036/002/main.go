package main

import (
	"io"
	"net/http"
)

type indexHandler int

func (h indexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "index")
}

type usersHandler int

func (h usersHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "user user user")
}

func main() {
	i := new(indexHandler)
	u := new(usersHandler)

	mux := http.NewServeMux()
	mux.Handle("/users/", u)
	mux.Handle("/", i)

	http.ListenAndServe(":8080", mux)
}
