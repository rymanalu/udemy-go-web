package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rymanalu/udemy-go-web/113/models"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Index</title>
		</head>
		<body>
			<p>
				<a href="/user/69">User Detail</a>
			</p>
		</body>
	</html>`

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "John Lennon",
		Gender: "male",
		Age:    40,
		ID:     p.ByName("id"),
	}

	uj, err := json.Marshal(u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uj)
}
