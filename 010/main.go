package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, `Release self-focus; embrace other focus.`)

	if err != nil {
		log.Fatalln(err)
	}
}
