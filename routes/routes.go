package routes

import (
	"golang-simple-http/controller"
	"net/http"
)

func LoadRoutes() {

	http.HandleFunc("/", controller.Index)
}
