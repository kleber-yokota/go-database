package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Create Category

	category := Category{Name: "Electronics"}
	db.Create(&category)

	categoryOffice := Category{Name: "Office"}
	db.Create(&categoryOffice)

	// Create Product
	product := []Product{{
		Name:       "Notebook",
		Price:      1000.0,
		CategoryID: category.ID,
	},
		{
			Name:       "Mouse",
			Price:      50.0,
			CategoryID: category.ID,
		},
		{
			Name:       "Paper",
			Price:      0.10,
			CategoryID: categoryOffice.ID,
		},
	}
	db.Create(&product)

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1,
	})

	db.Create(&SerialNumber{
		Number:    "7890",
		ProductID: 2,
	})

	db.Create(&SerialNumber{
		Number:    "102030",
		ProductID: 3,
	})

	// Get Products rows

	fmt.Println("simply display a row of table products")

	var products []Product
	db.Find(&products)
	for _, p := range products {
		fmt.Println(p.Name, p.Category.Name, p.SerialNumber.Number)
	}

	// Get Products with Category and Serial Number

	fmt.Println("Utilizing relational tables, display product information")

	var productsCategory []Product
	db.Preload("Category").Preload("SerialNumber").Find(&productsCategory)
	for _, p := range productsCategory {
		fmt.Println(p.Name, p.Category.Name, p.SerialNumber.Number)
	}

	// Get all products in each category
	fmt.Println("All items in the relevant category")
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			println("-", p.Name)
		}
	}

	// All serial numbers items in the relevant category
	fmt.Println("All serial numbers items in the relevant category")
	var categoriesSerialNumber []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categoriesSerialNumber).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categoriesSerialNumber {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			println("-", p.Name, "Serial Number:", p.SerialNumber.Number)
		}
	}

}
