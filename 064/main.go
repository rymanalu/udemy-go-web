package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", write)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func write(w http.ResponseWriter, r *http.Request) {
	incrementVisits(w, r)

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
}

func incrementVisits(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("visits")

	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "visits",
			Value: "0",
		}
	}

	counter, _ := strconv.Atoi(c.Value)
	counter++
	c.Value = strconv.Itoa(counter)

	http.SetCookie(w, c)
}
