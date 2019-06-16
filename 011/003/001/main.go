package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type artist struct {
	Name  string
	Quote string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	kurt := artist{
		Name:  "Kurt Cobain",
		Quote: "I'd rather be hated for who I am, than loved for who I am not.",
	}

	err := tpl.Execute(os.Stdout, kurt)

	if err != nil {
		log.Fatalln(err)
	}
}
