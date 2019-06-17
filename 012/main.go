package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type artist struct {
	Name  string
	Quote string
}

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
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

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", artists)

	if err != nil {
		log.Fatalln(err)
	}
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}
