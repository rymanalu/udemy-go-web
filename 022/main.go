package main

import (
	"bufio"
	"fmt"
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

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	defer conn.Close()

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
}
