package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rymanalu/udemy-go-web/115/models"
)

// UserController is a controller to handle things about user
type UserController struct{}

// New returns UserController
func New() *UserController {
	return &UserController{}
}

// Get returns a User by ID
func (uc UserController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "John Lennon",
		Gender: "male",
		Age:    40,
		ID:     p.ByName("id"),
	}

	uj, err := json.Marshal(u)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uj)
}

// Store stores a new User
func (uc UserController) Store(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u.ID = "007"

	uj, err := json.Marshal(u)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(uj)
}

// Destroy deletes a User by ID
func (uc UserController) Destroy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted"))
}
