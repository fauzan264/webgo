package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// route
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Home page"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, saya sedang belajar golang web"))
}
