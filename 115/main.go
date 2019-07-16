package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rymanalu/udemy-go-web/115/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.New()
	r.GET("/user/:id", uc.Get)
	r.POST("/user", uc.Store)
	r.DELETE("/user/:id", uc.Destroy)
	http.ListenAndServe(":8080", r)
}
