package model

import (
	database "golang-simple-http/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func FindAllProducts() []Product {
	db := database.ConnectWithDb()
	defer db.Close()
	dbProducts, err := db.Query("select * from product")
	p := Product{}
	products := []Product{}
	var id, amount int
	var name, description string
	var price float64
	if err != nil {
		panic(err.Error())
	}

	for dbProducts.Next() {
		err = dbProducts.Scan(&id, &name, &description, &price, &amount)
		p.Name = name
		p.Id = id
		p.Amount = amount
		p.Description = description
		p.Price = price
		products = append(products, p)
	}
	return products
}
