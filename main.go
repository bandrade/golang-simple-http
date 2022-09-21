package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func connectWithDb() *sql.DB {
	con := "user=admin dbname=market password=Passw0rd host=localhost sslmode=disable"
	db, err := sql.Open("postgres", con)
	if err != nil {
		panic((err.Error()))
	}
	return db
}

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

var temp = template.Must(template.ParseGlob("template/*.html"))

func main() {
	db := connectWithDb()
	defer db.Close()
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
