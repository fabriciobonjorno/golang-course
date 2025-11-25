package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           int     `gorm:"primaryKey"`
	Name         string  `gorm:"size:255"`
	Price        float64 `gorm:"type:decimal(10,2)"`
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
	gorm.Model
}

type SerialNumber struct {
	ID        int    `gorm:"primaryKey"`
	Number    string `gorm:"size:255"`
	ProductID int
}

func main() {
	// GORM: The fantastic ORM library for Golang, aims to be developer friendly.
	// https://gorm.io/

	// GORM 설치
	// go get -u gorm.io/gorm
	// go get -u gorm.io/driver/sqlite

	// GORM 문서
	// https://gorm.io/docs/

	// GORM GitHub
	// https://github.com/go-gorm/gorm

	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }

	// SQLite
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{Name: "Electronics"}
	db.Create(&category)

	products := []Product{
		{Name: "Mouse", Price: 29.99, CategoryID: category.ID},
		{Name: "Keyboard", Price: 49.99, CategoryID: category.ID},
		{Name: "Monitor", Price: 299.99, CategoryID: category.ID},
	}

	db.Create(&products)

	serialNumber := SerialNumber{Number: "SN123456789", ProductID: products[0].ID}
	db.Create(&serialNumber)

	// fetch product with its serial number
	var product Product
	db.Preload("SerialNumber").Preload("Category").First(&product, products[0].ID)
	println("Product:", product.Name, "Serial Number:", product.SerialNumber.Number, "Category:", product.Category.Name)

}
