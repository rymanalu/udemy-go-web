package main

import (
	"log"
	"os"
	"text/template"
)

type page struct {
	Title, Heading, Input string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	home := page{
		Title:   "Nothing Escaped",
		Heading: "Nothing is escaped with text/template",
		Input:   `<script>alert("BRUH");</script>`,
	}

	err := tpl.Execute(os.Stdout, home)

	if err != nil {
		log.Fatalln(err)
	}
}
