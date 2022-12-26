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
	fmt.Println("select one product")
	var product Product
	var productMouse Product

	db.First(&product, 1)
	fmt.Println(product)

	db.First(&productMouse, "name = ?", "Mouse")
	fmt.Println(productMouse)

	//select all
	fmt.Println("Select all")
	var productsFind []Product
	db.Find(&productsFind)
	for _, p := range productsFind {
		fmt.Println(p)
	}

	//select with limit
	fmt.Println("select with limits")
	var productLimit []Product
	db.Limit(2).Find(&productLimit)
	for _, pL := range productLimit {
		fmt.Println(pL)
	}

	//select with offset
	fmt.Println("select with offset")
	var productOffSet []Product
	db.Limit(2).Offset(2).Find(&productOffSet)
	for _, pL := range productOffSet {
		fmt.Println(pL)
	}

	//select with where
	fmt.Println("select with where")
	var productWhere []Product
	db.Where("price > ?", 90).Find(&productWhere)
	for _, pL := range productWhere {
		fmt.Println(pL)
	}

	//select with like
	fmt.Println("select with like")
	var productLike []Product
	db.Where("name like ?", "%book%").Find(&productLike)
	for _, pL := range productLike {
		fmt.Println(pL)
	}

}
