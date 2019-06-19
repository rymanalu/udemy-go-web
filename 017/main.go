package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 69
}

func (p person) AgeDouble() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p := person{
		Name: "John Lennon",
		Age:  40,
	}

	err := tpl.Execute(os.Stdout, p)

	if err != nil {
		log.Fatalln(err)
	}
}
