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

type guitar struct {
	Manufacturer string
	Model        string
}

type items struct {
	Performers  []artist
	Instruments []guitar
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

	f := guitar{
		Manufacturer: "Fender",
		Model:        "Stratocaster",
	}

	r := guitar{
		Manufacturer: "Rickenbacker",
		Model:        "325",
	}

	artists := []artist{kurt, john, dylan}
	guitars := []guitar{f, r}

	data := items{
		Performers:  artists,
		Instruments: guitars,
	}

	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatalln(err)
	}
}
