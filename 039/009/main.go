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

		strs := strings.Fields(str)

		body := fmt.Sprintf("Request Method: %s\nRequest URI: %s", strs[0], strs[1])
		fmt.Fprintf(conn, "HTTP/1.1 200OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
		fmt.Fprintf(conn, "\r\n")
		fmt.Fprintf(conn, body)

		i++
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
