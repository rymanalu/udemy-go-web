package main

import (
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	Email     string
	Password  []byte
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
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)

	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		e := r.FormValue("email")
		p := r.FormValue("password")
		fn := r.FormValue("first_name")
		ln := r.FormValue("last_name")

		if _, ok := usersDb[e]; ok {
			http.Error(w, "E-mail already taken, buddy", http.StatusUnprocessableEntity)
			return
		}

		sID, _ := uuid.NewV4()

		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		http.SetCookie(w, c)

		sessions[c.Value] = e

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		u := user{e, bs, fn, ln}

		usersDb[e] = u

		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
