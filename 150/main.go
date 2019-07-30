package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

type book struct {
	isbn   string
	title  string
	author string
	price  float32
}

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("postgres", "postgres://manalu:@localhost/udemy_bookstore?sslmode=disable")

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You connected to your database")
}

func main() {
	defer db.Close()

	router := httprouter.New()
	router.GET("/books", index)
	router.GET("/books/:isbn", show)
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

		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)

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

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}

func show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	isbn := p.ByName("isbn")

	if isbn == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	row := db.QueryRow("select * from books where isbn = $1", isbn)

	bk := book{}

	err := row.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)

	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
}
