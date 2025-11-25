package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int     `gorm:"primaryKey"`
	Name       string  `gorm:"size:255"`
	Price      float64 `gorm:"type:decimal(10,2)"`
	CategoryID int
	Category   Category
	gorm.Model
}

type Category struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"size:255"`
	Products []Product
	gorm.Model
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
	db.AutoMigrate(&Product{}, &Category{})

	category := Category{Name: "Electronics"}
	db.Create(&category)

	products := []Product{
		{Name: "Mouse", Price: 29.99, CategoryID: category.ID},
		{Name: "Keyboard", Price: 49.99, CategoryID: category.ID},
		{Name: "Monitor", Price: 299.99, CategoryID: category.ID},
	}

	db.Create(&products)

	// fetch product with its serial number
	var cat []Category
	err = db.Model(&Category{}).Preload("Products").Find(&cat).Error
	if err != nil {
		panic(err)
	}

	for _, c := range cat {
		println("Category:", c.Name)
		for _, p := range c.Products {
			println(" - Product:", p.Name, "Price:", p.Price)
		}
	}

}
