package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	handleError(err)

	defer li.Close()

	for {
		conn, err := li.Accept()

		handleError(err)

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	i := 1

	for scanner.Scan() {
		str := scanner.Text()

		if str == "" || i != 1 {
			break
		}

		var body string
		strs := strings.Fields(str)
		route := strs[0] + " " + strs[1]

		switch route {
		case "GET /apply":
			body = getApply()
		case "POST /apply":
			body = postApply()
		default:
			body = index()
		}

		fmt.Fprint(conn, "HTTP/1.1 200OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)

		i++
	}
}

func index() string {
	return `<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF=8">
		<title>Index - 011 - 039</title>
	</head>
	<body>
		<h1>Index</h1>
		<h3>Links</h3>
		<ul>
			<li><a href="/">Index</a></li>
			<li><a href="/apply">Apply</a></li>
		</ul>
	</body>
	</html>`
}

func getApply() string {
	return `<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF=8">
		<title>Apply - 011 - 039</title>
	</head>
	<body>
		<h1>Apply</h1>
		<h3>Links</h3>
		<ul>
			<li><a href="/">Index</a></li>
			<li><a href="/apply">Apply</a></li>
		</ul>
		<form action="/apply" method="post">
			<input type="submit" value="Apply">
		</form>
	</body>
	</html>`
}

func postApply() string {
	return `<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF=8">
		<title>Apply - 011 - 039</title>
	</head>
	<body>
		<h1>Apply</h1>
		<h3>Links</h3>
		<ul>
			<li><a href="/">Index</a></li>
			<li><a href="/apply">Apply</a></li>
		</ul>
		<p><strong>APPLIED!</strong></p>
	</body>
	</html>`
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
