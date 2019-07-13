package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type userIDContextKey string

type fNameContextKey string

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, userIDContextKey("userID"), 777)
	ctx = context.WithValue(ctx, fNameContextKey("fname"), "Bond")

	result := dbAccess(ctx)

	fmt.Fprintln(w, result)
}

func dbAccess(ctx context.Context) int {
	uid := ctx.Value(userIDContextKey("userID")).(int)

	return uid
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
