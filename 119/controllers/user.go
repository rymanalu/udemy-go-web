package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rymanalu/udemy-go-web/119/models"
	"gopkg.in/mgo.v2"
)

// UserController handles user
type UserController struct {
	session *mgo.Session
}

var users = make(map[string]models.User)

// NewUserController create a new UserController
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser show user by ID
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	u, ok := users[id]

	// Verify id is ObjectId hex representation, otherwise return status not found
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser handles creating user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	// u.ID = bson.NewObjectId().String()

	// store the user in mongodb
	users[u.ID] = u

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser deletes a user by ID
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if _, ok := users[id]; !ok {
		w.WriteHeader(404)
		return
	}

	// Delete user
	delete(users, id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}
