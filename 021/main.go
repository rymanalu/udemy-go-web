package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	defer li.Close()

	if err != nil {
		log.Panic(err)

		return
	}

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Println(err)

			continue
		}

		io.WriteString(conn, "\nHello from TCP server\n")

		fmt.Fprintln(conn, "How is your day?")

		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}
