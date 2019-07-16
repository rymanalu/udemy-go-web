package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
	"github.com/rymanalu/udemy-go-web/118/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.New(getSession())
	r.GET("/user", uc.Index)
	r.GET("/user/:id", uc.Show)
	r.POST("/user", uc.Store)
	r.DELETE("/user/:id", uc.Destroy)
	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s
}
