package routes

import (
	"golang-simple-http/controller"
	"net/http"
)

func LoadRoutes() {

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
}
