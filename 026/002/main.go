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

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Fatalln(err.Error())
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()

		fmt.Println(ln)

		if ln == "" {
			break
		}

		r := strings.Fields(ln)[1]

		body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Hello world!</title></head><body><h1>` + r + `</h1></body></html>`

		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	}
}
