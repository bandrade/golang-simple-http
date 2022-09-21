package controller

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"golang-simple-http/model"
)

var temp = template.Must(template.ParseGlob("template/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := model.FindAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		amount := r.FormValue("amount")
		price := r.FormValue("price")
		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error while converting price", err)
		}

		amountConverted, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error while converting price", err)
		}

		model.CreateNewProduct(name, description, priceConverted, amountConverted)

	}
	http.Redirect(w, r, "/", 301)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	model.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := model.FindProductById(id)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.URL.Query().Get("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		amount := r.FormValue("amount")
		price := r.FormValue("price")
		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error while converting price", err)
		}

		amountConverted, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error while converting price", err)
		}

		model.SaveProduct(id, name, description, priceConverted, amountConverted)

	}
	http.Redirect(w, r, "/", 301)
}
