package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/150.png", image)

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `<img src="/150.png">`)
}

func image(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "150.png")
}
