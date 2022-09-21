package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

var temp = template.Must(template.ParseGlob("template/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"T-SHIRT", "BLUE", 2.2, 1},
		{"Shoes", "BLUE", 2.2, 1},
	}
	temp.ExecuteTemplate(w, "Index", products)
}
