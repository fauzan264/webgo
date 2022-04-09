package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

// layout template
var viewLayout string = path.Join("views", "layout.html")

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	temp, err := template.ParseFiles(path.Join("views", "index.html"), viewLayout)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error: page not found", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":   "Home",
		"content": "Home page",
	}

	err = temp.Execute(w, data)
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

	temp, err := template.ParseFiles(path.Join("views", "product.html"), viewLayout)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error: page not found", http.StatusInternalServerError)
		return
	}

	// fmt.Fprintf(w, "Product page : %d", dataId)
	data := map[string]interface{}{
		"title": "Product",
		"content" : dataId,
	}
	err = temp.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error: page not found", http.StatusInternalServerError)
	}
}
