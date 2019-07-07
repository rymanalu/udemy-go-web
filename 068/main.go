package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	Email     string
	FirstName string
	LastName  string
}

var tpl *template.Template
var usersDb = map[string]user{}
var sessions = make(map[string]string)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")

	if err != nil {
		sID, _ := uuid.NewV4()

		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		http.SetCookie(w, c)
	}

	var u user
	if e, ok := sessions[c.Value]; ok {
		u = usersDb[e]
	}

	if r.Method == http.MethodPost {
		e := r.FormValue("email")
		fn := r.FormValue("first_name")
		ln := r.FormValue("last_name")

		u = user{e, fn, ln}

		sessions[c.Value] = e
		usersDb[e] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")

	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	e, ok := sessions[c.Value]

	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u := usersDb[e]

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
