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

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `<!DOCTYPE html>
	<html>
	<head>
		<title>Form</title>
	</head>
	<body>
		<form action="/" method="post">
			<input type="text" name="q">
			<input type="submit" value="Search">
		</form>
		<p>`+v+`</p>
	</body>
	</html>`)
}
