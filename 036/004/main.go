package main

import (
	"io"
	"net/http"
)

func i(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "index")
}

func u(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "user user user")
}

func main() {
	http.HandleFunc("/users/", u)
	http.HandleFunc("/", i)

	http.ListenAndServe(":8080", nil)
}
