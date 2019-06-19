package main

import (
	"log"
	"os"
	"text/template"
)

type fab struct {
	Name  string
	Quote string
	Alive bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	fab1 := fab{
		Name:  "John Lennon",
		Quote: "Life is what happens to you while you're busy making other plans.",
		Alive: false,
	}

	fab2 := fab{
		Name:  "Paul McCartney",
		Quote: "And in the end, the love you take is equal to the love you make.",
		Alive: true,
	}

	fab3 := fab{
		Name:  "George Harrison",
		Quote: "All the world is birthday cake, so take a piece, but not too much.",
		Alive: false,
	}

	fab4 := fab{
		Name:  "Ringo Starr",
		Quote: "Everything government touches turns to crap.",
		Alive: true,
	}

	fab5 := fab{
		Name:  "",
		Quote: "",
		Alive: false,
	}

	fabs := []fab{fab1, fab2, fab3, fab4, fab5}

	err := tpl.Execute(os.Stdout, fabs)

	if err != nil {
		log.Fatalln(err)
	}
}
