package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.gmao")

	if err != nil {
		log.Fatalln(err)
		return
	}

	err = tpl.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatalln(err)
		return
	}

	tpl, err = tpl.ParseFiles("two.gmao", "vespa.gmao")

	if err != nil {
		log.Fatalln(err)
		return
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)

	if err != nil {
		log.Fatalln(err)
		return
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)

	if err != nil {
		log.Fatalln(err)
		return
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)

	if err != nil {
		log.Fatalln(err)
		return
	}

	err = tpl.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatalln(err)
	}
}
