package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	li, err := net.Listen("tcp", ":8080")

	check(err)

	defer li.Close()

	for {
		conn, err := li.Accept()

		check(err)

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	i := 0

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		if i != 0 {
			continue
		}

		ln := scanner.Text()

		if ln == "" {
			continue
		}

		h := strings.Fields(ln)

		r := h[0] + " " + h[1]

		fmt.Println(r)

		switch r {
		case "GET /":
			sendResponse(conn, "index")
		default:
			sendResponse(conn, "404")
		}

		i++
	}
}

func sendResponse(conn net.Conn, name string) {
	f, err := os.Create("templates/" + name + ".html")
	defer f.Close()

	err = tpl.ExecuteTemplate(f, name+".gohtml", nil)

	check(err)

	h, err := os.Open("templates/" + name + ".html")
	defer h.Close()

	check(err)

	bs, err := ioutil.ReadAll(h)

	check(err)

	body := string(bs)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func check(e error) {
	if e != nil {
		log.Fatalln(e.Error())
	}
}
