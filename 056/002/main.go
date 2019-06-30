package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	var s string

	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		defer f.Close()

		bs, err := ioutil.ReadAll(f)

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		s = string(bs)

		dst, err := os.Create(filepath.Join("./storage/", h.Filename))

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		defer dst.Close()

		_, err = dst.Write(bs)

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `<!DOCTYPE html>
	<html>
	<head>
		<title>Upload</title>
	</head>
	<body>
		<form method="post" enctype="multipart/form-data">
			<input type="file" name="q">
			<input type="Submit">
		</form>
		<p>`+s+`</p>
	</body>
	</html>`)
}
