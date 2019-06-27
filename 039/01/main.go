package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Index Page")
	})

	http.HandleFunc("/dog/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Dog Page")
	})

	http.HandleFunc("/me/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "My Page")
	})

	http.ListenAndServe(":8080", nil)
}
