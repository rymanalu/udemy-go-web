package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

type book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

var db *sql.DB
var tpl *template.Template

func init() {
	var err error

	db, err = sql.Open("postgres", "postgres://username:password@localhost/database_name?sslmode=disable")

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You connected to your database")

	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	defer db.Close()

	router := httprouter.New()
	router.GET("/books", index)
	router.GET("/books/create", create)
	router.POST("/books/create", store)
	router.GET("/books/detail/:isbn", show)
	router.GET("/books/edit/:isbn", edit)
	router.POST("/books/edit/:isbn", update)
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := db.Query("select * from books")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	bks := make([]book, 0)

	for rows.Next() {
		bk := book{}

		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "index.gohtml", bks)
}

func show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	isbn := p.ByName("isbn")

	if isbn == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	row := db.QueryRow("select * from books where isbn = $1", isbn)

	bk := book{}

	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)

	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
}

func create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "create.gohtml", nil)
}

func store(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bk := book{
		Isbn:   r.FormValue("isbn"),
		Title:  r.FormValue("title"),
		Author: r.FormValue("author"),
	}

	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

	f64, err := strconv.ParseFloat(p, 32)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

	bk.Price = float32(f64)

	_, err = db.Exec("insert into books (isbn, title, author, price) values ($1, $2, $3, $4)", bk.Isbn, bk.Title, bk.Author, bk.Price)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

func edit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	isbn := p.ByName("isbn")

	if isbn == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	row := db.QueryRow("select * from books where isbn = $1", isbn)

	bk := book{}

	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)

	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "edit.gohtml", bk)
}

func update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	isbn := p.ByName("isbn")

	if isbn == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	row := db.QueryRow("select * from books where isbn = $1", isbn)

	bk := book{}

	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)

	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	bk = book{
		Isbn:   r.FormValue("isbn"),
		Title:  r.FormValue("title"),
		Author: r.FormValue("author"),
	}

	pr := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || pr == "" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

	f64, err := strconv.ParseFloat(pr, 32)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

	bk.Price = float32(f64)

	_, err = db.Exec("update books set isbn = $1, title = $2, author = $3, price = $4 where isbn = $1", bk.Isbn, bk.Title, bk.Author, bk.Price)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
