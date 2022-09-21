package controller

import (
	"net/http"
	"text/template"

	"golang-simple-http/model"
)

var temp = template.Must(template.ParseGlob("template/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := model.FindAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}
