package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)

	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")

		if err != nil {
			log.Fatalln(err)
		}

		defer mf.Close()

		fs := strings.Split(fh.Filename, ".")

		ext := fs[len(fs)-1]

		h := sha1.New()

		io.Copy(h, mf)

		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		wd, err := os.Getwd()

		if err != nil {
			log.Fatalln(err)
		}

		path := filepath.Join(wd, "public", "pics", fname)

		nf, err := os.Create(path)

		if err != nil {
			log.Fatalln(err)
		}

		defer nf.Close()

		mf.Seek(0, 0)

		io.Copy(nf, mf)

		c = appendValue(w, c, fname)
	}

	http.SetCookie(w, c)

	var xs []string

	if len(c.Value) > 0 {
		xs = strings.Split(c.Value, "|")
	}

	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")

	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: "",
		}
	}

	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	var xs []string

	if len(c.Value) > 0 {
		xs = strings.Split(c.Value, "|")
	}

	if !contains(xs, fname) {
		xs = append(xs, fname)
	}

	c.Value = strings.Join(xs, "|")

	return c
}

func contains(xs []string, name string) bool {
	for _, v := range xs {
		if v == name {
			return true
		}
	}

	return false
}
