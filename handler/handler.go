package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"webgo/entity"
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
	// id := r.URL.Query().Get("id")

	// dataId, err := strconv.Atoi(id)

	// if err != nil || dataId < 1 {
	// 	http.NotFound(w, r)
	// 	return
	// }

	temp, err := template.ParseFiles(path.Join("views", "product.html"), viewLayout)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error: page not found", http.StatusInternalServerError)
		return
	}

	// fmt.Fprintf(w, "Product page : %d", dataId)
	dataProduct := []entity.Product{
		{Id: 1, Name: "Kaos Polos Hitam", Price: 40000, Stock: 20},
		{Id: 2, Name: "Kaos Polos Putih", Price: 40000, Stock: 10},
		{Id: 3, Name: "Kaos Polos Biru", Price: 40000, Stock: 25},
		{Id: 4, Name: "Baju Polo", Price: 90000, Stock: 5},
		{Id: 5, Name: "Celana Chino", Price: 100000, Stock: 1},
		{Id: 6, Name: "Hoodie", Price: 140000, Stock: 29},
		{Id: 7, Name: "Celana Jeans", Price: 240000, Stock: 22},
		{Id: 8, Name: "Kaos Polos Merah", Price: 40000, Stock: 9},
		{Id: 9, Name: "Kaos Polos Cream", Price: 40000, Stock: 18},
	}
	data := map[string]interface{}{
		"title":   "Product",
		"content": dataProduct,
	}
	err = temp.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error: page not found", http.StatusInternalServerError)
	}
}

// learn http method
func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("It's a GET"))
	case "POST":
		w.Write([]byte("It's a POST"))
	default:
		http.Error(w, "Error: bad request", http.StatusBadRequest)
	}
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error: The form must be filled", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		email := r.Form.Get("email")
		message := r.Form.Get("message")

		// result data contact
		dataContact := map[string]string{
			"Name":    name,
			"Email":   email,
			"Message": message,
		}

		// data
		data := map[string]interface{}{
			"title":   "Result form contact",
			"content": dataContact,
		}

		// menampilkan html page
		temp, err := template.ParseFiles(path.Join("views", "result.html"), viewLayout)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error: page not found", http.StatusInternalServerError)
			return
		}

		// menampilkan data
		err = temp.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error: internal server error", http.StatusInternalServerError)
			return
		}
		return

	}

	http.Error(w, "Error: bad request", http.StatusBadRequest)
}
