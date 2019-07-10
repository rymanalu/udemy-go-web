package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/udemy_go_web?charset=utf8mb4")

	checkError(err)

	defer db.Close()

	err = db.Ping()

	checkError(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.icon", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)

	checkError(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Successfully completed.")

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
