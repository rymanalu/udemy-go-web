package main

import (
	"log"
	"os"
	"text/template"
)

type menu struct {
	Breakfast []string
	Lunch     []string
	Dinner    []string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := menu{
		Breakfast: []string{"Kopi", "Susu", "Roti", "Telur"},
		Lunch:     []string{"Ayam Goreng", "Nasi Cakalang"},
		Dinner:    []string{"Nasi Goreng", "Sate"},
	}

	err := tpl.Execute(os.Stdout, m)

	if err != nil {
		log.Fatalln(err)
	}
}
