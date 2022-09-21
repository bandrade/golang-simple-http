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

func CreateNewProduct(name, description string, price float64, amount int) {
	db := database.ConnectWithDb()
	defer db.Close()

	insert, err := db.Prepare("insert into product(name,description, price, amount) values ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(name, description, price, amount)

}
