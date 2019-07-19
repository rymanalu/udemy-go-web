package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rymanalu/udemy-go-web/120/controllers"
	"github.com/rymanalu/udemy-go-web/120/models"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	session := make(map[string]models.User)

	f, err := os.Open("users.json")

	if err != nil {
		return session
	}

	defer f.Close()

	err = json.NewDecoder(f).Decode(&session)

	if err != nil {
		return session
	}

	return session
}
