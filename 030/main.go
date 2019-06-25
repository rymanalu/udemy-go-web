package main

import (
	"fmt"
	"net/http"
)

type web int

func (w web) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello world!")
}

func main() {
	w := new(web)

	http.ListenAndServe(":8080", w)
}
