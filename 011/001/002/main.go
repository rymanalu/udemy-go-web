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
	fabs := []string{"John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr"}

	err := tpl.Execute(os.Stdout, fabs)

	if err != nil {
		log.Fatalln(err)
	}
}
