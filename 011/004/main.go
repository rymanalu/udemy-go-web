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

	john := artist{
		Name:  "John Lennon",
		Quote: "Life is what happens to you while you're busy making other plans.",
	}

	dylan := artist{
		Name:  "Bob Dylan",
		Quote: "How many roads must a man walk down Before your can call him a man? . . . The answer, my friend, is blowin' in the wind, The answer is blowin' in the wind.",
	}

	artists := []artist{kurt, john, dylan}

	err := tpl.Execute(os.Stdout, artists)

	if err != nil {
		log.Fatalln(err)
	}
}
