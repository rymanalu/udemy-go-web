package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
	"github.com/rymanalu/udemy-go-web/118/models"
)

type users []models.User

// UserController is a controller to handle things about user
type UserController struct {
	session *mgo.Session
}

// New returns UserController
func New(s *mgo.Session) *UserController {
	return &UserController{s}
}

// Index shows all users
func (uc UserController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	xu := new(users)

	if err := uc.session.DB("udemy-go-web").C("users").Find(nil).All(xu); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	xuj, err := json.Marshal(xu)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", xuj)
}

// Show returns a User by ID
func (uc UserController) Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("udemy-go-web").C("users").FindId(oid).One(&u); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
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

	u.ID = bson.NewObjectId()

	err = uc.session.DB("udemy-go-web").C("users").Insert(u)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(u)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

// Destroy deletes a User by ID
func (uc UserController) Destroy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("udemy-go-web").C("users").RemoveId(oid); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Deleted user", oid)
}
