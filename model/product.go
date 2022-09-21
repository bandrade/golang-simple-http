package model

import (
	database "github.com/bandrade/golang-simple-http/db"
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

func DeleteProduct(id string) {
	db := database.ConnectWithDb()
	defer db.Close()

	delete, err := db.Prepare("delete from product where id=$1")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)

}

func FindProductById(idToEdit string) Product {
	db := database.ConnectWithDb()
	dbProducts, err := db.Query("select * from product where id =" + idToEdit)
	p := Product{}
	var id, amount int
	var name, description string
	var price float64
	if err != nil {
		panic(err.Error())
	}

	dbProducts.Next()
	err = dbProducts.Scan(&id, &name, &description, &price, &amount)
	p.Name = name
	p.Id = id
	p.Amount = amount
	p.Description = description
	p.Price = price
	return p

}

func SaveProduct(id, name, description string, price float64, amount int) {
	db := database.ConnectWithDb()
	defer db.Close()

	edit, err := db.Prepare("update product set name=$1,description=$2,price=$3,amount=$4 where id=$5)")
	if err != nil {
		panic(err.Error())
	}
	edit.Exec(name, description, price, amount, id)

}
