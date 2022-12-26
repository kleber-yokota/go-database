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

	// update row
	var p Product
	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(&p)

	//get updated row
	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)

	// delete row
	db.Delete(&p2)

	//verify deleted row
	var p3 Product
	db.First(&p3, 1)
	fmt.Println(p3.Name)
}
