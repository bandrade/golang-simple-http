package main

import (
	"net/http"

	"github.com/bandrade/golang-simple-http/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
