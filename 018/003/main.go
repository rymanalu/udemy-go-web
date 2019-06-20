package main

import (
	"log"
	"os"
	"text/template"
)

type restaurant struct {
	Breakfast []string
	Lunch     []string
	Dinner    []string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := restaurant{
		Breakfast: []string{"Kopi", "Susu", "Roti", "Telur"},
		Lunch:     []string{"Ayam Goreng", "Nasi Cakalang"},
		Dinner:    []string{"Nasi Goreng", "Sate"},
	}

	err := tpl.Execute(os.Stdout, r)

	if err != nil {
		log.Fatalln(err)
	}
}
