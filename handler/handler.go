package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("Home page"))
	temp, err := template.ParseFiles(path.Join("views", "index.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error: page not found", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error: page not found", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, saya sedang belajar golang web"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	dataId, err := strconv.Atoi(id)

	if err != nil || dataId < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Product page : %d", dataId)
}
