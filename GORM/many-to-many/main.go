package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// Create Category

	category := Category{Name: "Electronics"}
	db.Create(&category)

	categoryOffice := Category{Name: "Office"}
	db.Create(&categoryOffice)

	// Create Product
	product := []Product{
		{
			Name:       "Notebook",
			Price:      1000.0,
			Categories: []Category{category, categoryOffice},
		},
		{
			Name:       "Mouse",
			Price:      50.0,
			Categories: []Category{category},
		},
		{
			Name:       "Paper",
			Price:      0.10,
			Categories: []Category{categoryOffice},
		},
	}
	db.Create(&product)

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

}
