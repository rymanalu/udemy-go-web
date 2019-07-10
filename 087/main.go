package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:g3m5c00l@tcp(127.0.0.1:3306)/udemy_go_web?charset=utf8mb4")

	checkError(err)

	defer db.Close()

	err = db.Ping()

	checkError(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.icon", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)

	checkError(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintln(w, "at index")

	checkError(err)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`select name from amigos;`)

	checkError(err)

	defer rows.Close()

	var s, name string

	for rows.Next() {
		err = rows.Scan(&name)

		checkError(err)

		s += name + "\n"
	}

	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`create table if not exists customers(id int auto_increment primary key, name varchar(255) not null);`)

	checkError(err)

	defer stmt.Close()

	r, err := stmt.Exec()

	checkError(err)

	n, err := r.RowsAffected()

	checkError(err)

	fmt.Fprintln(w, "Customers table created", n)
}

func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`insert into customers (name) values ("John");`)

	checkError(err)

	defer stmt.Close()

	r, err := stmt.Exec()

	checkError(err)

	n, err := r.RowsAffected()

	checkError(err)

	fmt.Fprintln(w, "Rows Affected:", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("select * from customers;")

	checkError(err)

	defer rows.Close()

	var id int
	var name string

	for rows.Next() {
		err = rows.Scan(&id, &name)

		checkError(err)

		fmt.Printf("ID: %d, NAME: %s\n", id, name)

		fmt.Fprintln(w, "RETRIEVED RECORD:", id, name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`update customers set name = "Paul" where name = "John";`)

	checkError(err)

	defer stmt.Close()

	r, err := stmt.Exec()

	checkError(err)

	n, err := r.RowsAffected()

	checkError(err)

	fmt.Fprintln(w, "UPDATED RECORD:", n)
}

func del(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`delete from customers where name = "Paul";`)

	checkError(err)

	defer stmt.Close()

	r, err := stmt.Exec()

	checkError(err)

	n, err := r.RowsAffected()

	checkError(err)

	fmt.Fprintln(w, "DELETED RECORD:", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare("drop table if exists customers;")

	checkError(err)

	defer stmt.Close()

	_, err = stmt.Exec()

	checkError(err)

	fmt.Fprintln(w, "Customers table dropped")
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
