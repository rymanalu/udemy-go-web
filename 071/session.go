package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
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

	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")

	if err != nil {
		return false
	}

	e := sessions[c.Value]

	_, ok := usersDb[e]

	return ok
}
