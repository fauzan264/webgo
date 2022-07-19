package main

import (
	"log"
	"net/http"
	"webgo/handler"
)

func main() {
	mux := http.NewServeMux()

	// route
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/process", handler.Process)

	// assets
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting web on port 8100")

	err := http.ListenAndServe(":8100", mux)
	log.Fatal(err)
}
