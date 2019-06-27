package main

import (
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	handleError(err)

	defer li.Close()

	for {
		conn, err := li.Accept()

		handleError(err)

		io.WriteString(conn, "I see you connected.\n")

		conn.Close()
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
