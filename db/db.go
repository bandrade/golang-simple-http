package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectWithDb() *sql.DB {
	con := "user=admin dbname=market password=Passw0rd host=localhost sslmode=disable"
	db, err := sql.Open("postgres", con)
	if err != nil {
		panic((err.Error()))
	}
	return db
}
