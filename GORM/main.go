package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create batch
	products := []Product{
		{Name: "Notebook", Price: 1000.0},
		{Name: "Mouse", Price: 50.0},
		{Name: "Keyboard", Price: 100.0},
	}
	db.Create(&products)

	//select one

	var product Product
	var productMouse Product

	db.First(&product, 1)
	fmt.Println(product)

	db.First(&productMouse, "name = ?", "Mouse")
	fmt.Println(productMouse)

	//select all

	var productsFind []Product
	db.Find(&productsFind)
	for _, p := range productsFind {
		fmt.Println(p)
	}

}
