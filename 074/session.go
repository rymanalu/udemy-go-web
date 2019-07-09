package main

import (
	"net/http"
	"time"

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
	}

	c.MaxAge = sessionLength

	http.SetCookie(w, c)

	var u user

	if s, ok := sessions[c.Value]; ok {
		u = usersDb[s.e]
	}

	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")

	if err != nil {
		return false
	}

	s := sessions[c.Value]

	_, ok := usersDb[s.e]

	return ok
}

func cleanSessions() {
	for k, v := range sessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * sessionLength) {
			delete(sessions, k)
		}
	}
}
