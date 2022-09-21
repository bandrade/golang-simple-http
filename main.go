package main

import (
	"net/http"

	"golang-simple-http/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
