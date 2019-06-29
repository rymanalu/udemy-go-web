package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/image", image)

	http.ListenAndServe(":8080", nil)
}

func image(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `<img src="/150.png">`)
}
