package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session-id")

	if err != nil {
		id, _ := uuid.NewV4()

		c := &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			HttpOnly: true,
		}

		http.SetCookie(w, c)
	}

	fmt.Println(c)
}
