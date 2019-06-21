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

type restaurant struct {
	Name    string
	Address string
	Menu    menu
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

	r := restaurant{
		Name:    "Warung Burjo",
		Address: "Mabes",
		Menu:    m,
	}

	err := tpl.Execute(os.Stdout, r)

	if err != nil {
		log.Fatalln(err)
	}
}
