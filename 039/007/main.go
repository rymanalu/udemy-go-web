package main

import (
	"bufio"
	"fmt"
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

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		str := scanner.Text()

		if str == "" {
			break
		}

		fmt.Println(str)
	}

	fmt.Println("Code got here.")

	io.WriteString(conn, "I see you connected.\n")
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
