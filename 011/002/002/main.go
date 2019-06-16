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
	// fabs := []string{"John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr"}

	fabs := map[string]string{
		"Rhythm Guitar": "John Lennon",
		"Bass":          "Paul McCartney",
		"Lead Guitar":   "George Harrison",
		"Drum":          "Ringo Starr",
	}

	err := tpl.Execute(os.Stdout, fabs)

	if err != nil {
		log.Fatalln(err)
	}
}
