package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int     `gorm:"primaryKey"`
	Name  string  `gorm:"size:255"`
	Price float64 `gorm:"type:decimal(10,2)"`
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

	dsn := "root:root@tcp(localhost:3306)/goexpert"
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
	db.AutoMigrate(&Product{})

	product := Product{
		Name:  "Notebook",
		Price: 1999.99,
	}

	db.Create(&product)

	products := []Product{
		{Name: "Mouse", Price: 29.99},
		{Name: "Keyboard", Price: 49.99},
		{Name: "Monitor", Price: 299.99},
	}

	db.Create(&products)
}
