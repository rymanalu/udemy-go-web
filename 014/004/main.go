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
	game := struct {
		Score1 int
		Score2 int
	}{
		6,
		9,
	}

	err := tpl.Execute(os.Stdout, game)

	if err != nil {
		log.Fatalln(err)
	}
}
