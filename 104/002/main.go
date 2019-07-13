package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
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

	result, err := dbAccess(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, result)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	defer cancel()

	ch := make(chan int)

	go func() {
		// Example of long running task...
		uid := ctx.Value(userIDContextKey("userID")).(int)

		time.Sleep(5 * time.Second)

		if ctx.Err() != nil {
			return
		}

		ch <- uid
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
