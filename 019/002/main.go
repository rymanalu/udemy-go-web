package main

import (
	"html/template"
	"log"
	"os"
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
		Title:   "Escaped",
		Heading: "Danger is escaped with html/template",
		Input:   `<script>alert("BRUH");</script>`,
	}

	err := tpl.Execute(os.Stdout, home)

	if err != nil {
		log.Fatalln(err)
	}
}
