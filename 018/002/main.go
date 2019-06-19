package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		hotel{"Mandapa", "Jl. Kedewatan, Banjar Kedewatan", "Ubud", "Bali"},
		hotel{"Raffles", "Jl. Prof.Dr Satrio 3-5", "Jakarta Selatan", "DKI Jakarta"},
		hotel{"Tugu", "Sigar Pejalin", "Lombok", "Nusa Tenggara Barat"},
	}

	err := tpl.Execute(os.Stdout, hotels)

	if err != nil {
		log.Fatalln(err)
	}
}
