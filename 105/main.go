package main

import "net/http"

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}
