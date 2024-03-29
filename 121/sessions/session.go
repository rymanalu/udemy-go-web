package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rymanalu/udemy-go-web/121/models"
	uuid "github.com/satori/go.uuid"
)

// Length is session length
const Length int = 30

// Users is users db
var Users = map[string]models.User{} // user ID, user
// Sessions is sessions db
var Sessions = map[string]models.Session{} // session ID, session
// LastCleaned is last time the session cleaned
var LastCleaned = time.Now()

// GetUser returns user based on session
func GetUser(w http.ResponseWriter, req *http.Request) models.User {
	// get cookie
	ck, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		ck = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	ck.MaxAge = Length
	http.SetCookie(w, ck)

	// if the user exists already, get user
	var u models.User
	if s, ok := Sessions[ck.Value]; ok {
		s.LastActivity = time.Now()
		Sessions[ck.Value] = s
		u = Users[s.UserName]
	}
	return u
}

// AlreadyLoggedIn checks if user already logged in
func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	ck, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := Sessions[ck.Value]
	if ok {
		s.LastActivity = time.Now()
		Sessions[ck.Value] = s
	}
	_, ok = Users[s.UserName]
	// refresh session
	ck.MaxAge = Length
	http.SetCookie(w, ck)
	return ok
}

// Clean cleans session
func Clean() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	Show()                      // for demonstration purposes
	for k, v := range Sessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(Sessions, k)
		}
	}
	LastCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	Show()                     // for demonstration purposes
}

// Show shows session for demonstration purposes
func Show() {
	fmt.Println("********")
	for k, v := range Sessions {
		fmt.Println(k, v.UserName)
	}
	fmt.Println("")
}
