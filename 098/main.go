package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	c = appendValue(w, c)
	http.SetCookie(w, c)
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")

	if err != nil {
		sID, err := uuid.NewV4()

		if err != nil {
			log.Fatalln(err)
		}

		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}

	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	// Values...
	p1 := "john.jpg"
	p2 := "paul.jpg"
	p3 := "george.jpg"
	p4 := "ringo.jpg"

	// Append...
	s := c.Value

	if !strings.Contains(s, p1) {
		s += "|" + p1
	}

	if !strings.Contains(s, p2) {
		s += "|" + p2
	}

	if !strings.Contains(s, p3) {
		s += "|" + p3
	}

	if !strings.Contains(s, p4) {
		s += "|" + p4
	}

	c.Value = s

	return c
}
